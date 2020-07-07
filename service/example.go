package service

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"github.com/zhuxuyang/grpc_example/protos"
)

type Example struct {
}

func (s *Example) Hello(ctx context.Context, req *protos.HelloRequest) (resp *protos.HelloResponse, err error) {
	return &protos.HelloResponse{
		Answer: viper.GetString("name"),
		Time:   time.Now().Unix(),
	}, nil
}
