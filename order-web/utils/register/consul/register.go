package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Registry struct {
	Host string
	Port int
}

type RegistryClient interface {
	Register(address string, port int, name string, tags []string, id string) error
	DeRegister(serviceId string) error
}

func NewRegistryClient(host string, port int) RegistryClient {
	return &Registry{
		Host: host,
		Port: port,
	}
}

func DNClient(Host string, Port int) (*api.Client, error) {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", Host, Port)

	client, err := api.NewClient(cfg)
	return client, err
}

func (r *Registry) DeRegister(serviceId string) error {
	client, err := DNClient(r.Host, r.Port)
	if err != nil {
		return err
	}
	err = client.Agent().ServiceDeregister(serviceId)
	return err
}

func (r *Registry) Register(address string, port int, name string, tags []string, id string) error {
	//cfg := api.DefaultConfig()
	//cfg.Address = fmt.Sprintf("%s:%d", r.Host, r.Port)
	//
	//client, err := api.NewClient(cfg)
	client, err := DNClient(r.Host, r.Port)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
		Timeout:                        "37s",
		Interval:                       "37s",
		DeregisterCriticalServiceAfter: "260s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}
