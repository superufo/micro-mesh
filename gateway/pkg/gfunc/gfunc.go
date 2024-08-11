package gfunc

import (
	"context"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

// 全局执行需要的方法
var (
	CFG = g.Cfg()
)

func init() {
	scfgs := make([]constant.ServerConfig, 0)
	ccfg := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
	)

	ctx := context.Background()
	nacosAddr, _ := gcfg.Instance().Get(ctx, "nacos.Address")
	for _, addr := range nacosAddr.Strings() {
		scfg := constant.ServerConfig{}
		addStr := strings.Split(addr, ":")

		scfg.IpAddr = addStr[0]
		port, _ := strconv.Atoi(addStr[1])
		scfg.Port = uint64(port)
		scfg.ContextPath = "/nacos"
		scfgs = append(scfgs, scfg)
	}

	ns := NewNacosService(ctx, g.Log(), scfgs, ccfg)

	Nacos = *ns
}

// GetConfigFromNacos 服务配置
func GetConfigFromNacos(ctx context.Context) {
	mode, err := gcfg.Instance().Get(ctx, "gf.gmode")
	if err != nil {
		g.Log().Fatalf(ctx, "获取配置错误:", err)
	}
	group, err := gcfg.Instance().Get(ctx, "nacos.group")
	if err != nil {
		g.Log().Fatalf(ctx, "获取配置错误:", err)
	}

	g.Log().Info(ctx, mode.String(), group.String())
	//服务发现
	config, err := Nacos.GetConfig(mode.String(), group.String())
	CFG.GetAdapter().(*gcfg.AdapterFile).SetContent(config, "default.yaml")
	CFG.GetAdapter().(*gcfg.AdapterFile).SetFileName("default.yaml")
}

func GetSerAddr(proto string, path string) string {
	srvName := GetSerName(proto, path)
	url, _ := Nacos.GetNamingClient().GetService(vo.GetServiceParam{
		ServiceName: srvName,
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		//GroupName:   "group-a",             // 默认值DEFAULT_GROUP
	})

	return url
}

// 根据后端协议,前端路径获取后端服务名称
func GetSerName(proto string, path string) string {
	ctx := context.Background()
	p, err := g.Cfg().Get(ctx, "micro.grpc")
	if err != nil {
		g.Log().Fatalf(ctx, "配置中没有配置微服务设置", err)
	}

	for _, i := range p.Array() {
		ip := gconv.Map(i)
		mapUrl := strings.Split(ip["mapUrl"].(string), ",")
		for _, u := range mapUrl {
			if u == path {
				return ip["name"].(string)
			}
		}
	}

	return ""
}

// 限流  sentinel.rules.api
func limitBySentinel(ctx context.Context, ruleStr string) bool {
	rules := make([]*flow.Rule, 1)
	rulesVar, _ := gcfg.Instance().Get(ctx, ruleStr)
	g.Log().Info(ctx, "配置:", rulesVar)
	if err := gconv.Struct(rulesVar.Map(), &rules[0]); err != nil {
		g.Log().Info(ctx, "配置转换出错:", rulesVar)
	}

	isLimit := NewSentinelService(ctx, g.Log(), rules).Limit()

	return isLimit
}

func HttpLimit(ctx context.Context) {

}

func TcpLimit(ctx context.Context) {

}

func UdpLimit(ctx context.Context) {

}

func WebsocketLimit(ctx context.Context) {

}
