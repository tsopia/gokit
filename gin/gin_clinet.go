package gin

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)
type Client struct {
	engine *gin.Engine
	srv    *http.Server
}

func NewClient(addr string) *Client {
	r := gin.Default()
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &Client{
		engine: r,
		srv:    srv,
	}
}

func (c *Client) RegisterRoutes(routes func(r *gin.Engine)) {

	routes(c.engine)
}
func (c *Client) Use(middleware func(c *gin.Context)) {
	c.engine.Use(middleware)
}
func (c *Client) Run() {
	go func() {
		// 服务连接
		if err := c.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := c.srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
