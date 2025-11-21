// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"blogV2/blog-rpc/blog"
	"blogV2/gateway-api/internal/config"
	"blogV2/user-rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC usercenter.UserCenterClient
	BlogRPC blog.BlogClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: usercenter.NewUserCenterClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
		BlogRPC: blog.NewBlogClient(zrpc.MustNewClient(c.BlogRpcConf).Conn()),
	}
}
