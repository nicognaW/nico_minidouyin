package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	config2 "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"log"
	"nico_minidouyin/config"
)

func CreateConsulRegistry() (r *registry.Registry) {
	// build a consul client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.ConsulAddress
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// build a consul register with the consul client
	consulRegister := consul.NewConsulRegister(consulClient)
	r = &consulRegister
	return
}

func WithConsul() (optPtr *config2.Option) {
	if DevEnv() {
		optPtr = &config2.Option{F: func(o *config2.Options) {}}
		return
	}
	opt := server.WithRegistry(*CreateConsulRegistry(), &registry.Info{
		ServiceName: config.FeedServiceName,
		Addr:        utils.NewNetAddr("tcp", config.ServiceAddress),
		Weight:      10,
		Tags:        nil,
	})
	optPtr = &opt
	return
}
