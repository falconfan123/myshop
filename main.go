package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"myshop/internal/cmd"
	_ "myshop/internal/logic"
	_ "myshop/internal/packed" //
)

func main() {
	cmd.Main.Run(gctx.New())
}
