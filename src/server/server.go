package server

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
)

var rpcClient client.Client

func init() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 初始化服务
	service := micro.NewService(
		micro.Version("latest"),
		micro.Registry(etcdReg),
	)

	rpcClient = service.Client()
}

func ReferClient() client.Client {
	return rpcClient
}
