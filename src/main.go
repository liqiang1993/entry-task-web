package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/routers"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	webService := web.NewService(
		web.Name("phoenix_gateway"),
		web.Address("127.0.0.1:8001"),
		web.Handler(routers.InitRouter()),
		web.Registry(etcdReg))

	err := webService.Init()
	if err != nil {
		log.Error("server init failed, err:%s", err)
		return
	}

	log.Info("ready start http server listening %s", config.GetGlobalConfig().ServerSetting.HttpPort)

	err = webService.Run()
	if err != nil {
		log.Error("server run failed, err:%s", err)
		return
	}

	log.Info("start http server listening %s", config.GetGlobalConfig().ServerSetting.HttpPort)
}
