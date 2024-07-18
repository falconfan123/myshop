package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Rotation, // 轮播图
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
