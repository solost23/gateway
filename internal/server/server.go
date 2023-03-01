package server

import (
	"fmt"
	"gateway/configs"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	sc *configs.ServerConfig
}

func NewServer(sc *configs.ServerConfig) *Server {
	return &Server{
		sc: sc,
	}
}

func (s *Server) Run() {
	router := gin.Default()

	for _, route := range s.sc.Routes {
		router.GET(route.Path, func(ctx *gin.Context) {
			proxy := &httputil.ReverseProxy{
				Director: func(req *http.Request) {
					req.URL.Scheme = route.Scheme
					req.URL.Host = route.Host
				},
			}
			proxy.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	go func() {
		if err := router.Run(s.sc.Addr); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			// 服务停止
			// 服务注销
			fmt.Println("服务注销成功")
			// kafka.ConsumerClient.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
