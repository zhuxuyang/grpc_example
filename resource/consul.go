package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type ConsulResource struct {
	consulClient *api.Client
	serverID     string
	serverIP     string
}

var CR *ConsulResource
var UnHealthyChan chan bool

// getConsulResource returns the singleton of consul
func getConsulResource() (*ConsulResource, error) {
	token := viper.GetString("golang-consul.token")
	address := viper.GetString("golang-consul.host")
	p := viper.GetString("golang-consul.port")
	if token == "" || address == "" || p == "" {
		err := errors.New(fmt.Sprintf(
			"consul configuration empty [host: %v][port: %v][token: %v]",
			address, p, token))
		log.Println(err)
		return nil, err
	}

	consulConfig := &api.Config{
		Address: fmt.Sprintf("%s:%s", address, p),
		Token:   token,
	}
	client, err := api.NewClient(consulConfig)
	if err != nil {
		err := errors.New(fmt.Sprintf("NewClient consul error\t%v \t address : %s", err, address))

		return nil, err
	}

	ServerIP := os.Getenv("HOST")
	if ServerIP == "" {
		err := errors.New("HOST should be defined in system environment" +
			" and it's value should be inner IP in K8S cluster")
		log.Println(err)
		return nil, err
	}

	return &ConsulResource{consulClient: client,
		serverID: fmt.Sprintf("%s-%s", viper.GetString("service"), ServerIP),
		serverIP: ServerIP,
	}, nil
}

// RegisterService 注册服务
func RegisterConsulService() {
	log.Println(viper.GetString("env"))
	if viper.GetString("env") == "" || viper.GetString("env") == "dev" {
		return
	}
	UnHealthyChan = make(chan bool)
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		select {
		case <-UnHealthyChan:
			writer.WriteHeader(500)
			_, _ = writer.Write([]byte("bye"))
		default:
			writer.WriteHeader(200)
			_, _ = writer.Write([]byte("ok"))
		}
	})

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("healthPort")), nil)
		if err != nil {
			Logger.Errorf("start health failed, err is %s", err.Error())
		}
	}()

	consulResource, err := getConsulResource()
	CR = consulResource
	if err != nil || consulResource == nil {
		log.Panic(fmt.Sprintf("RegisterConsulService err %v  %v", consulResource, err))
	}

	ServerGRPCPortString := viper.GetString("port")
	ServerGRPCPortInt, err := strconv.ParseInt(ServerGRPCPortString, 10, 64)
	if err != nil {
		log.Panic("port config in yaml should be integer instead of" + ServerGRPCPortString)
	}
	HealthyPortHTTP := viper.GetString("healthPort")
	if HealthyPortHTTP == "" {
		log.Panic("healthPort config in yaml is empty instead of" + ServerGRPCPortString)
	}

	//register consul
	agent := CR.consulClient.Agent()
	interval := time.Duration(3) * time.Second
	deregister := time.Duration(1) * time.Second
	reg := &api.AgentServiceRegistration{
		ID:      CR.serverID,                     // 服务节点的名称
		Name:    viper.GetString("service"),      // 服务名称
		Tags:    []string{"go", "advertisement"}, // tag，可以为空
		Port:    int(ServerGRPCPortInt),          // 服务端口
		Address: CR.serverIP,                     // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       interval.String(), // 健康检查间隔
			HTTP:                           fmt.Sprintf("http://%s:%s/health", CR.serverIP, HealthyPortHTTP),
			DeregisterCriticalServiceAfter: deregister.String(), // 注销时间，相当于过期时间
		},
	}

	s, _ := json.MarshalIndent(&reg, "", "   ")
	log.Println(string(s))
	err = agent.ServiceRegister(reg)
	log.Println(fmt.Sprintf("agent.ServiceRegister %v", err))
	if err != nil {
		log.Panic(fmt.Sprintf("Service Register error : %v", err))
	}
}

// DeregisterService sign off from consul by defined serviceID
func (cr *ConsulResource) DeregisterService() error {
	if viper.GetString("env") == "" || viper.GetString("env") == "dev" {
		return nil
	}
	agent := cr.consulClient.Agent()
	if agent == nil {
		return errors.New("fail get consul client agent()")
	}
	return cr.consulClient.Agent().ServiceDeregister(cr.serverID)
}
