package example

import (
	"context"
	"fmt"
	"github.com/tsopia/gokit/log"
	"testing"
)

func TestLogger(t *testing.T) {
	//config := &log.LoggerConfig{
	//	Level: "debug",
	//}
	//log.InitLogger(config)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value1")
	ctx = context.WithValue(ctx, "trace", "value2")
	log.Info(ctx).Msg("sss")

	log.Error(ctx).Ctx(ctx).Err(fmt.Errorf("erreasaaaa")).Msg("sas")

}
