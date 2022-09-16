package server

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
)

var rpcClient client.Client

func init() {
	service := micro.NewService()
	service.Init()
	rpcClient = service.Client()
}

func ReferClient() client.Client {
	return rpcClient
}
