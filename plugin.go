package main

import (
	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/arg"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 控制程序
	Binary string `default:"${BINARY=kubectl}"`
	// 服务
	Endpoint string `default:"ENDPOINT" validate:"required"`
	// 命名空间
	Namespace string `default:"${NAMESPACE=default}"`

	// 用户名
	Username string `default:"${USERNAME=default}"`
	// 密钥
	Token string `default:"${TOKEN}" validate:"required"`

	// 版本
	Version string `default:"${VERSION=v1}"`
	// 名称
	Name string `default:"${NAME}" validate:"required"`
	// 注册表
	Registry string `default:"${REGISTRY}" validate:"required"`
	// 仓库
	Repository string `default:"${REPOSITORY}" validate:"required"`

	// 无状态服务
	Deployment *_deployment `default:"${DEPLOYMENT}"`
	// 服务配置
	Service *service `default:"${SERVICE}"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		drone.NewStep(newSetupStep(p)).Name("配置").Build(),
		drone.NewStep(newDeploymentStep(p)).Name("无状态服务").Build(),
	}
}

func (p *plugin) Fields() (fields gox.Fields[any]) {
	fields = make(gox.Fields[any], 0, 2)
	if nil != p.Deployment {
		fields = append(fields, field.New("deployment", p.Deployment))
	}
	if nil != p.Service {
		fields = append(fields, field.New("service", p.Service))
	}

	return
}

func (p *plugin) kubectl(args arg.Args) error {
	return p.Command(p.Binary).Args(args...).Build().Exec()
}
