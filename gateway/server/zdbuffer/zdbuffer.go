package zdbuffer

import (
	proto  "zhugedaojia.com/common/net/pb/gw"
)

type zdbuffer interface {

}

// 使用chan可以缓存数据
type HttpBuffer struct {
     TosData      chan 	map[string][]proto.UnitaryTos                      // string  每个url对应要发给后端的数据
     TogData      chan  map[string][]proto.UnitaryTog                      // 后端微服务回复给网关的数据
}

func NewHttpBuffer() *HttpBuffer {
	buffer := HttpBuffer{
		TogData: make(chan 	map[string][]proto.UnitaryTog,1000),
		TosData: make(chan 	map[string][]proto.UnitaryTos,1000),
	}

	return &buffer
}








