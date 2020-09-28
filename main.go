package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zhuxuyang/grpc_example/config"
	"github.com/zhuxuyang/grpc_example/grpc"
	"github.com/zhuxuyang/grpc_example/resource"
)

type A interface {
	Do()
}

type B interface {
	FF()
}

type S struct {
	B string
}

func (a S) FF() {
	log.Println("ff", a.B)
}
func (a S) Do() {
	log.Println("do", a.B)
}

func main() {
	log.Println("service_example start...")

	config.InitConfig()
	resource.InitLogger()
	//resource.InitDB()
	//
	grpc.RegisterGRPCService()
	resource.RegisterConsulService()
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stopSignal:
		resource.CR.DeregisterService()
	}
}
