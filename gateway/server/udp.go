package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/micro-mesh/gateway/server/base"
	"github.com/micro-mesh/gateway/server/inf"
	"time"
)

type UdpSrv struct {
	// 数据
	Ctx context.Context

	Host       string
	Port       int
	properties UdpProperties

	base.BaseServer
}

func (h UdpSrv) Run(ops inf.Properties) {
	h.Ctx = context.Background()

	g.Log().Info(h.Ctx, fmt.Sprint("UdpSrv ~~~~~~~~~~~   %s:%d", h.Host, h.Port))
	gudp.NewServer(fmt.Sprint("%s:%d", h.Host, h.Port), func(conn *gudp.Conn) {
		defer conn.Close()
		for {
			if data, _ := conn.Recv(-1); len(data) > 0 {
				fmt.Println(string(data))
			}
		}
	}).Run()

	g.Log().Info(h.Ctx, fmt.Sprint("UdpSrv %s:%d", h.Host, h.Port))
}

func (h UdpSrv) GetProperties() inf.Properties {
	return h.properties
}

func (h UdpSrv) SetProperties(properties inf.Properties) {
	h.properties = properties.(UdpProperties)
}

func NewUdpSrv(host string, port int) inf.IServer {
	updSrv := UdpSrv{
		Ctx:        gctx.New(),
		Host:       host,
		Port:       port,
		properties: UdpProperties{},
	}

	updSrv.Name = "udp"
	return updSrv
}

func WithUdpTimeout(timeout time.Duration) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(UdpProperties)
		hp.timeout = timeout
	}
}

func WithUdpCash(caching bool) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(UdpProperties)
		hp.caching = caching
	}
}

type UdpProperties struct {
	timeout time.Duration
	caching bool
}

func (h UdpProperties) Get(key string) (val interface{}) {
	return nil
}
