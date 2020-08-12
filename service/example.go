package service

import (
	"context"
	"google.golang.org/grpc/metadata"
	"time"

	"github.com/spf13/viper"
	"github.com/zhuxuyang/grpc_example/protos"
)

type Example struct {
}

func (s *Example) Hello(ctx context.Context, req *protos.HelloRequest) (resp *protos.HelloResponse, err error) {
	//id := GetRequestIdFromContext(ctx)
	//resource.Logger.Info("id=?", id)
	//m := model.UserA{}
	//err = resource.Db.
	//	Model(&model.UserA{}).Where("1=1").
	//	Scan(&m).Error
	////resource.Logger.Info("query: err ", err)
	//id = GetRequestIdFromContext(ctx)
	//resource.Logger.Info("id=?", id)

	return &protos.HelloResponse{
		Answer: viper.GetString("name"),
		Time:   time.Now().Unix(),
	}, nil
}

func GetRequestIdFromContext(c context.Context) string {
	if t, ok := metadata.FromIncomingContext(c); ok {
		data := t.Get("request_id")
		if len(data) > 0 {
			return data[0]
		}
	}
	return ""
}
