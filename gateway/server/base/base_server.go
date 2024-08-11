package base

import (
	"context"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
	"github.com/micro-mesh/gateway/server/inf"

	"github.com/gogf/gf/v2/net/ghttp"
)

type BaseServer struct {
	Name string

	IsClose bool
	Ctx     context.Context
	Host    string
	Port    int

	//properties inf.Properties
}

func (h BaseServer) Receive() []byte {
	return nil
}

func (h BaseServer) Send([]byte) error {
	return nil
}

func (h BaseServer) Run(ops inf.Properties) {}

func (h BaseServer) Close() {}

func (h BaseServer) GetName() string {
	return h.Name
}

func (h BaseServer) SetName(name string) {
	h.Name = name
}

//func (h BaseServer) GetProperties() inf.Properties {
//	return h.properties
//}
//
//func (h BaseServer) SetProperties(properties inf.Properties){
//	h.properties = properties
//}

type HttpHandle func(group *ghttp.RouterGroup)
type TcpHandle func(*gtcp.Conn)
type UdpHandle func(*gudp.Conn)
