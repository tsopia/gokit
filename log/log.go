package log

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// 获取当前文件的绝对路径
	_, b, _, _ = runtime.Caller(0)
	// 获取当前文件的目录
	basepath = filepath.Dir(b)
)

var logger = zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(2).Logger()

func init() {
	// 自定义CallerMarshalFunc
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		// 将绝对路径转换为相对路径
		relativePath, err := filepath.Rel(basepath, file)
		if err != nil {
			return file
		}
		// 移除路径前面的"../"
		relativePath = strings.TrimPrefix(relativePath, "../")
		return fmt.Sprintf("%s:%d", relativePath, line)
	}
}

func Info(ctx context.Context) *zerolog.Event {
	event := logger.Info()
	// 从context.Context中获取"trace"的值
	if traceValue := ctx.Value("trace"); traceValue != nil {
		event = event.Str("trace", traceValue.(string))
	}

	// 输出日志
	return event
}

func Error(ctx context.Context) *zerolog.Event {
	event := logger.Error()
	// 从context.Context中获取"trace"的值
	if traceValue := ctx.Value("trace"); traceValue != nil {
		event = event.Str("trace", traceValue.(string))
	}

	// 输出日志
	return event
}
