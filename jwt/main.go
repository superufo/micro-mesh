package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/micro-mesh/jwt/internal/cmd"
)

func main() {
	index := &gcmd.Command{
		Name:        "main",
		Brief:       "start http server",
		Description: "this is the command entry for starting your process",
	}

	err := index.AddCommand(&cmd.HttpCmd, &cmd.GrpcCmd)
	if err != nil {
		g.Log().Error("err:", err)
	}

	index.Run(gctx.New())
}
