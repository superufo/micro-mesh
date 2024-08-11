package inf

type  IServer interface {
	Receive() []byte
	Send([]byte) error
	Run(ops Properties)
	GetName() string
	SetName(name string)

	GetProperties() Properties
	SetProperties(properties Properties)

	Close()
}

// 属性接口
type Properties interface {
    Get(key string) (val interface{})
}

type Property interface {
	apply(*Properties)
}

//optionFunc  继承 Option
type PropertyFunc func(*Properties)

// apply 执行 optionFunc 自身
func (f PropertyFunc) apply(o *Properties) {
	f(o)
}







