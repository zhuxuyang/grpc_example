package main

import (
	"log"

	"github.com/zhuxuyang/grpc_example/config"
	"github.com/zhuxuyang/grpc_example/grpc"
)

func main() {
	log.Println("service_example start...")
	config.InitConfig()
	grpc.RegisterGRPCService()

	select {}
}
