package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/rpc"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/setting"
	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/util"
	"github.com/lucky-cheerful-man/phoenix_gateway/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化配置
	setting.InitConfig()

	// 初始化日志
	log.InitLog()

	// 初始化工具
	util.InitUtil()

	// 初始化rpc服务
	rpc.InitRPC()

	// 信号处理
	dealSignal()

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: setting.AppSetting.MaxHeaderBytes,
	}

	log.Infof("start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Warnf("server init failed, err:%s", err)
	}
}

func dealSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		for s := range sigs {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT:
				log.Warnf("got signal:%v and try to exit: ", s)
				os.Exit(0)
			default:
				log.Warnf("other signal:%v: ", s)
			}
		}
	}()
}
