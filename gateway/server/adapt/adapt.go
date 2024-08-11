package adapt

type BaseAdapt interface {
	Send() error   // 代理端转发给微服务端
	Reply() ([]byte,error)   // 微服务端回复代理端信息

    BroadCastReply()([]byte,error)
	GroupReply()([]byte,error)
}

var ProtoAdapt = &protoAdapt{}

type protoAdapt struct {
	f    string      // 前端协议
	fbt  int         // 前端与后端的编解码 1 proto3  2 json
	b    string      // 后端协议

	fbd  []byte      // 前端给后端的数据
	bfd  []byte      // 后端回复给前端的数据
}

// 代理端转发给微服务端
func (pa protoAdapt)Send() error{
	return nil
}

// 微服务端 代理端
func (pa protoAdapt)Reply() ([]byte,error) {
	return nil,nil
}

func (pa protoAdapt)BroadCastReply() ([]byte,error) {
	return nil,nil
}

func (pa protoAdapt)GroupReply() ([]byte,error) {
	return nil,nil
}













