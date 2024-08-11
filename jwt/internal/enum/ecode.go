package enum

import "github.com/gogf/gf/v2/errors/gcode"

var (
	RecordNotFindCode = gcode.New(-1000, "信息不存在", "detailed description")
)

func GetNewCode(code gcode.Code, desc string) gcode.Code {
	return gcode.New(code.Code(), code.Message(), desc)
}
