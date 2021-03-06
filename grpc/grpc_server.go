package grpc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zhuxuyang/grpc_example/protos"
	"github.com/zhuxuyang/grpc_example/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
)

var grpc_log_file = "./grpc.log"

func getRecoverOption() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		start := bytes.Index(buf, []byte("/src/runtime/panic.go"))
		end := bytes.Index(buf, []byte("\ngoroutine "))
		if len(buf) > end && end > start {
			buf = buf[start:end]
		} else {
			buf = buf[0:stackSize]
		}

		log.Println("grpc panic \n", p, string(buf))
		return nil
	})
}

func RegisterGRPCService() {
	gRPCLogDecider := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		return true
	}
	logger := logrus.New()
	fileName := path.Join(viper.GetString("logDir"), "rpc-course.log")
	// Create a LogFile
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Printf("create logfile if failed, err is %v", err)
	}
	logger.SetFormatter(&MyGrpcLogFormat{})
	logger.Out = logFile
	logrusEntry := logrus.NewEntry(logger)

	logOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logrusEntry, logOpts...),
			grpc_logrus.PayloadUnaryServerInterceptor(logrusEntry, gRPCLogDecider),
			grpc_recovery.UnaryServerInterceptor(getRecoverOption()),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logrusEntry, logOpts...),
			grpc_logrus.PayloadStreamServerInterceptor(logrusEntry, gRPCLogDecider),
			grpc_recovery.StreamServerInterceptor(getRecoverOption()),
		),
	)
	// 注册服务
	protos.RegisterExampleServer(grpcServer, &service.Example{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("grpc_port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}()

	log.Println("GRPC注册成功 port:", viper.GetString("grpc_port"))
}

type MyGrpcLogFormat struct{}

func (f *MyGrpcLogFormat) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	b.WriteString(entry.Time.Format("[2006-01-02 15:04:05] "))
	b.WriteString(fmt.Sprintf("[%v] ", strings.ToUpper(entry.Level.String())))
	b.WriteString(fmt.Sprintf("[%v ", entry.Data))
	b.WriteByte('\n')

	return b.Bytes(), nil
}
