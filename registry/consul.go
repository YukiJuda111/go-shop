package registry

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

type consul struct {
	client *api.Client
}

var Reg Register

// 确保consul结构体实现了对应的接口
var _ Register = (*consul)(nil)

// Init 连接至consul服务，初始化全局的consul对象
func Init(addr string) (err error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	Reg = &consul{c}
	return
}

// RegisterService 将gRPC服务注册到consul
func (c *consul) RegisterService(serviceName string, ip string, port int, tags []string) error {
	// 健康检查
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port), // 这里一定是外部可以访问的地址
		Timeout:                        "10s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "20s",
	}
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    tags,                                           // 为服务打标签
		Address: ip,
		Port:    port,
		Check:   check,
	}
	return c.client.Agent().ServiceRegister(srv)
}

// ListService 服务发现
func (c *consul) ListService(serviceName string) (map[string]*api.AgentService, error) {
	return c.client.Agent().ServicesWithFilter(fmt.Sprintf("Service==`%s`", serviceName))
}

// Deregister 注销服务
func (c *consul) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}
