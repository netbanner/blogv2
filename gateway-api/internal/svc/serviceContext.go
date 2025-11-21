// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"blogV2/gateway-api/internal/config"
	"blogV2/user-rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC usercenter.UserCenter
	BlogRPC zrpc.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpcConf)),
		BlogRPC: zrpc.MustNewClient(c.BlogRpcConf),
	}
}
