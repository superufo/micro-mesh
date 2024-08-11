package cmd

import (
	"context"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"

	"github.com/micro-mesh/jwt/internal/controller/user"
)

var (
	GrpcCmd = gcmd.Command{
		Name:  "grpc",
		Usage: "grpc",
		Brief: "start grpc server of jwt",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			rules := make([]*flow.Rule, 1)
			rulesVar, err := gcfg.Instance().Get(ctx, "sentinel.rules.api")
			g.Log().Info(ctx, "配置:", rulesVar)
			if err := gconv.Struct(rulesVar.Map(), &rules[0]); err != nil {
				g.Log().Info(ctx, "配置转换出错:", rulesVar)
			}

			isLimit := NewSentinelService(ctx, g.Log(), rules).Limit()
			if isLimit == true {
				os.Exit(0)
			}

			wg := sync.WaitGroup{}
			wg.Add(1)
			func() {
				var (
					scfgs = make([]constant.ServerConfig, 0)
					ccfg  = *constant.NewClientConfig(
						constant.WithNamespaceId(""),
						constant.WithTimeoutMs(5000),
						constant.WithNotLoadCacheAtStart(true),
						constant.WithLogDir("./log/"),
						constant.WithCacheDir("./manifest/"),
						constant.WithLogLevel("debug"),
					)

					nacosAddr = &gvar.Var{}
				)

				nacosAddr, err = gcfg.Instance().Get(ctx, "nacos.Address")
				for _, addr := range nacosAddr.Strings() {
					scfg := constant.ServerConfig{}
					addStr := strings.Split(addr, ":")

					scfg.IpAddr = addStr[0]
					port, _ := strconv.Atoi(addStr[1])
					scfg.Port = uint64(port)
					scfg.ContextPath = "/nacos"
					scfgs = append(scfgs, scfg)
				}
				nacosService := NewNacosService(ctx, g.Log(), scfgs, ccfg)

				// 配置下载
				gmode, err := gcfg.Instance().Get(ctx, "gf.gmode")
				if err != nil {
					g.Log().Fatalf(ctx, "获取配置错误:", err)
				}
				group, err := gcfg.Instance().Get(ctx, "nacos.group")
				if err != nil {
					g.Log().Fatalf(ctx, "获取配置错误:", err)
				}

				g.Log().Info(ctx, gmode.String(), group.String())
				//服务发现
				config, err := nacosService.GetConfig(gmode.String(), group.String())
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetContent(config, "default.yaml")
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("default.yaml")
				g.Log().Info(ctx, "config:", g.Cfg())

				// 服务注册
				add, err := gcfg.Instance().Get(ctx, "server.grpc.Address")
				name, err := gcfg.Instance().Get(ctx, "server.grpc.Name")
				addStr := strings.Split(add.String(), ":")
				port, _ := strconv.Atoi(addStr[1])

				nacosService.Register(name.String(), addStr[0], uint64(port))
				g.Log().Info(ctx, "ctx:", ctx, "scfgs:", scfgs, "name:", name.String(), "addStr:", addStr, "port:", port)

				wg.Done()
			}()

			add, err := gcfg.Instance().Get(ctx, "server.grpc.Address")
			name, err := gcfg.Instance().Get(ctx, "server.grpc.Name")
			conf := &grpcx.GrpcServerConfig{
				Address:          add.String(),
				Name:             name.String(),
				Logger:           g.Log(),
				LogPath:          "",
				LogStdout:        true,
				ErrorStack:       true,
				ErrorLogEnabled:  true,
				ErrorLogPattern:  "",
				AccessLogEnabled: false,
				AccessLogPattern: "",
				Options:          nil,
			}
			s := grpcx.Server.New(conf)
			user.Register(s, ctx)
			s.Run()

			g.Log().Info(ctx, "s:", s)
			return nil
		},
	}
)
