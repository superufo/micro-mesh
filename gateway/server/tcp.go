package server

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
	"zhugedaojia.com/gateway/server/base"
	"zhugedaojia.com/gateway/server/inf"
)

type  TcpSrv struct {
	// 数据
	Ctx context.Context
	Host    string
	Port    int
	properties TcpProperties

	base.BaseServer
}

func (h TcpSrv)Receive() []byte{
	return nil
}

func (h TcpSrv)Send([]byte) error {
	return nil
}

func (h TcpSrv)Run(ops inf.Properties) {
	h.Ctx = context.Background()
	g.Log().Info(h.Ctx,fmt.Sprint("tcp ~~~ %s:%d",h.Host,h.Port))

	gtcp.NewServer(fmt.Sprint("%s:%d",h.Host,h.Port), func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}

			if err != nil {
				break
			}
		}
	}).Run()

	g.Log().Info(h.Ctx,fmt.Sprint("tcp ~~~ %s:%d",h.Host,h.Port))
}

func (h TcpSrv) GetProperties() inf.Properties {
	return h.properties
}

func (h TcpSrv) SetProperties(properties inf.Properties){
	h.properties = properties.(TcpProperties)
}

func NewTcpSrv(host string,port int) inf.IServer {
	tcpSrv := TcpSrv{
		Ctx:  gctx.New(),
		Host: host,
		Port: port,
		properties:  TcpProperties{},
	}

	tcpSrv.Name = "tcp"
	return tcpSrv
}

func  WithTcpTimeout(timeout  time.Duration) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(TcpProperties)
		hp.timeout = timeout
	}
}

func WithTcpCash(caching  bool) inf.PropertyFunc {
	return func(p *inf.Properties) {
		hp := (*p).(TcpProperties)
		hp.caching = caching
	}
}

type TcpProperties struct {
	timeout time.Duration
	caching bool
}

func (h  TcpProperties)Get(key string) (val interface{}){
	return nil
}


