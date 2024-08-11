package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/micro-mesh/jwt/internal/cmd"
)

func main() {
	cmd.GrpcCmd.Run(gctx.New())
}
