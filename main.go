package main

import (
	"github.com/zhuxuyang/grpc_example/config"
	"github.com/zhuxuyang/grpc_example/grpc"
	"github.com/zhuxuyang/grpc_example/resource"
	"log"
)

func main() {
	log.Println("service_example start...")
	config.InitConfig()
	resource.InitLogger()
	//resource.InitDB()

	grpc.RegisterGRPCService()

	select {}
}
