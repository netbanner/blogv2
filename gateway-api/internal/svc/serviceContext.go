// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"blogV2/blog-rpc/blog"
	"blogV2/blog-rpc/blogClient"
	"blogV2/gateway-api/internal/config"
	"blogV2/gateway-api/internal/middleware"
	"blogV2/user-rpc/userCenterClient"
	"blogV2/user-rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRPC       userCenterClient.UserCenter
	BlogRPC       blogClient.Blog
	JwtMiddleware *middleware.JwtMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRPC:       usercenter.NewUserCenterClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
		BlogRPC:       blog.NewBlogClient(zrpc.MustNewClient(c.BlogRpcConf).Conn()),
		JwtMiddleware: middleware.NewJwtMiddleware(c.Auth.AccessSecret),
	}
}
