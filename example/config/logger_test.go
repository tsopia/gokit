package example

import (
	"context"
	"fmt"
	"github.com/tsopia/gokit/log"
	"testing"
)

func TestLogger(t *testing.T) {
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
	//logger := zerolog.New(os.Stderr)
	log.InitLog()
	// 添加请求ID和用户ID到logger
	ctx := log.With().
		Str("request_id", "123").
		Str("user_id", "abc").
		Logger().WithContext(context.Background())
	// 创建一个context
	//ctx := context.Background()

	// 使用WithContext方法将logger与context关联起来
	//ctx = logger.WithContext(ctx)

	// 在后续的代码中，我们可以通过context获取到logger

	log.Info(ctx).Msg("info")
	log.Debug(ctx).Msg("Debug")
	log.Warn(ctx).Msg("Warn")
	log.Error(ctx).Err(fmt.Errorf("errorsssss")).Msg("error")
	log.Fatal(ctx).Err(fmt.Errorf("fatalerrorsssss")).Msg("fatal")
	log.Panic(ctx).Err(fmt.Errorf("panicerrorsssss")).Msg("panic")

}

//func f() {
//	logger := zerolog.New(os.Stdout)
//	ctx := context.Background()
//	// Attach the Logger to the context.Context
//	ctx = logger.WithContext(ctx)
//	ctx = context.WithValue(ctx, "key", "value")
//
//	someFunc(ctx)
//}
//
//func someFunc(ctx context.Context) {
//	// Get Logger from the go Context. if it's nil, then
//	// `zerolog.DefaultContextLogger` is returned, if
//	// `DefaultContextLogger` is nil, then a disabled logger is returned.
//	logger := zerolog.Ctx(ctx)
//	logger.Info().Ctx(ctx).Str("key", ctx.Value("key")).Msg("Hello")
//}
