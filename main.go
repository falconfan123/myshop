package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"myshop/internal/cmd"
	_ "myshop/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New()) //调用internal/cmd包的对应命令引导程序启动
	//gctx.New表示创建一个带链路跟踪特性的context上下文对象给下游链路
}
