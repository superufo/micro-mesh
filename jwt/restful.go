package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"
	"zhugedaojia.com/jwt/internal/cmd"
)

func main() {
	cmd.HttpCmd.Run(gctx.New())
}

