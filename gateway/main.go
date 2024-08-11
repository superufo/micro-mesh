package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/micro-mesh/gateway/pkg/gfunc"
	"github.com/micro-mesh/gateway/server"
	"github.com/micro-mesh/gateway/server/inf"
)

var (
	ctx     = gctx.New()
	srvList = make([]inf.IServer, 0)
	opsMap  = make(map[string][]inf.PropertyFunc, 0)

	propertiesMap = make(map[string]inf.Properties, 0)
)

func main() {
	go gfunc.GetConfigFromNacos(ctx)

	proxy, err := gcfg.Instance().Get(ctx, "gateway.proxy")
	g.Log().Info(ctx, "proxy,err:", proxy, err)

	l := proxy.Map()
	g.Log().Info(ctx, "proxy:", l["tcp"])
	if l["udp"] != "" {
		host, port := getHp(l["udp"].(string))
		udpSrv := server.NewUdpSrv(host, port)
		udpSrv.SetProperties(server.UdpProperties{})
		srvList = append(srvList, udpSrv)

		// 需要设置的属性
		opsMap["udp"] = append(opsMap["udp"], server.WithUdpTimeout(5*time.Second))
	}

	if l["tcp"] != "" {
		host, port := getHp(l["tcp"].(string))
		tcpSrv := server.NewUdpSrv(host, port)
		srvList = append(srvList, tcpSrv)

		g.Log().Info(ctx, "tcp", "host:", host, "port:", port)
	}

	if l["websocket"] != "" {
		host, port := getHp(l["websocket"].(string))
		wsSrv := server.NewWsSrv(host, port)
		srvList = append(srvList, wsSrv)
	}

	if l["http"] != "" {
		host, port := getHp(l["http"].(string))
		httpSrv := server.NewHttpSrv(host, port)
		srvList = append(srvList, httpSrv)

		timeout := server.WithHttpTimeout(5 * time.Second)
		ops := make([]inf.PropertyFunc, 0)
		opsMap["http"] = append(ops, timeout)
	}

	g.Log().Info(ctx, "srvList:", srvList)

	for _, srv := range srvList {
		//var properties  inf.Properties
		//if p, ok := propertiesMap[name]; ok {
		//	properties = p
		//}
		name := srv.GetName()
		properties := srv.GetProperties()
		g.Log().Info(ctx, "opsMap[name],len", opsMap[name], len(opsMap[name]))

		if opsMap[name] != nil && len(opsMap[name]) > 0 {
			for i, f := range opsMap[name] {
				g.Log().Info(ctx, i)
				f(&properties)
			}
		}

		g.Log().Info(ctx, "name, properties:", name, properties)
		srv.Run(properties)
	}

	select {}
}

func getHp(addr string) (host string, port int) {
	var err error
	arr := strings.Split(addr, ":")
	host = arr[0]
	port, err = strconv.Atoi(arr[1])
	if err != nil {
		g.Log().Errorf(ctx, "解析地址错误")
	}

	return host, port
}
