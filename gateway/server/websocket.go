package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"time"
	"zhugedaojia.com/gateway/server/base"
	"zhugedaojia.com/gateway/server/inf"
)

type  WsSrv struct {
	// 数据
	Ctx context.Context
	Srv   *ghttp.Server

	Host    string
	Port    int
	properties  WsProperties

	base.BaseServer
}

func (h WsSrv)Receive() []byte{
	return nil
}

func (h WsSrv)Send([]byte) error {
	return nil
}

func (h WsSrv)Run(ops inf.Properties) {
	h.Ctx = context.Background()
	s := g.Server()

	s.BindHandler("/ws", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(h.Ctx, err)
			r.Exit()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			if err = ws.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	s.SetServerRoot(gfile.MainPkgPath())

	g.Log().Info(h.Ctx,fmt.Sprint("ws ~~~ %s:%d",h.Host,h.Port))
	s.SetPort(h.Port)
	s.Run()
}

func (h WsSrv) GetProperties() inf.Properties {
	return h.properties
}

func (h WsSrv) SetProperties(properties inf.Properties){
	h.properties = properties.(WsProperties)
}

func NewWsSrv(host string,port int) inf.IServer {
	wsSrv := WsSrv{
		Ctx:  gctx.New(),
		Host: host,
		Port: port,
		properties: WsProperties{},
	}

	wsSrv.Name = "websocket"
	return wsSrv
}

func  WithWSTimeout(timeout  time.Duration) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(WsProperties)
		hp.timeout = timeout
	}
}

func WithWsCash(caching  bool) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(WsProperties)
		hp.caching = caching
	}
}

type WsProperties struct {
	timeout time.Duration
	caching bool
}

func (h  WsProperties)Get(key string) (val interface{}){
	return nil
}


