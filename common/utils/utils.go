/*
 * @Author: yanghang
 * @Date: 2021-08-26 10:02:57
 * @LastEditors: yanghang
 * @Description:
 */
package utils

import (
	"crypto/md5"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/os/gtime"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/crypto/bcrypt"
)

// 解析markdown为html
func MarkdownToHtml(mdContent string) string {
	return string(blackfriday.Run([]byte(mdContent)))
}

func Md5Str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

//密码hash加密
func PasswordHash(passwd string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), cost)
	return string(bytes), err
}

//密码hash解密
func PasswordVerify(passwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
	return err == nil
}

//将from结构体指针的字段赋值到to指针结构体的字段
func StructAssign(from interface{}, to interface{}) {
	fromVal := reflect.ValueOf(from).Elem()
	toVal := reflect.ValueOf(to).Elem()
	fromTypeOfT := fromVal.Type()

	for i := 0; i < fromVal.NumField(); i++ {
		name := fromTypeOfT.Field(i).Name
		if ok := toVal.FieldByName(name).IsValid(); ok {
			toVal.FieldByName(name).Set(reflect.ValueOf(fromVal.Field(i).Interface()))
		}
	}
}

//从 id=>text 的map中 根据以“，”拼接的文本中获取文本供展示用
func GetTextFromTextMapByIds(ids string, textMap *map[string]string) (text string) {
	if ids == "" {
		return
	}
	idStrList := strings.Split(ids, ",")
	textList := []string{}
	for _, v := range idStrList {
		if t, ok := (*textMap)[v]; ok {
			textList = append(textList, t)
		}
	}
	text = strings.Join(textList, ",")
	return
}

func GetZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func GetAge(birthday *gtime.Time) (age int) {
	age = time.Now().Year() - birthday.Year()

	if int(time.Now().Month()) < birthday.Month() {
		age--
	}

	if age < 1 {
		age = 1
	}
	return
}

//将Float64转成指定位数的Float64
func Float(value float64, precision int) (res float64) {
	n10 := math.Pow10(precision)
	return math.Trunc((value+0.5/n10)*n10) / n10
}

//比较浮点数，比较浮点数，当bigValue>smallValue返回1，当bigValue<smallValue返回-1，当相等时返回0
func FloatCompare(bigValue float64, smallValue float64) (res int) {
	n10 := math.Pow10(10)
	zoomBigValue := math.Trunc((bigValue + 0.5/n10) * n10)
	zoomSmallValue := math.Trunc((smallValue + 0.5/n10) * n10)
	if zoomBigValue > zoomSmallValue {
		return 1
	} else if zoomBigValue < zoomSmallValue {
		return -1
	}
	return 0
}

//格式化Float64显示
func FormatFloat(value float64, precision int) (res string) {
	return fmt.Sprintf("%."+strconv.Itoa(precision)+"f", Float(value, precision))
}

func InviteCode(uid uint64) (code string) {
	const (
		BASE    = "PULYVWXGHO72B8QRCEIJKAZ956TDMNFS" //这个不能改，改了可能会重复的情况了
		DECIMAL = 32
		PAD     = "3" // 0  1  3  4  差位补全用  0和1不用,容易和I,O字母混淆   4 不够吉利,那就先默认用3吧
		LEN     = 8
	)
	id := uid
	mod := uint64(0)
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		code += string(BASE[mod])
	}
	codeLen := len(code)
	if codeLen < LEN {
		code += PAD
		for i := 0; i < LEN-codeLen-1; i++ {
			code += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return
}

// 手机号中间4位替换为*号
func FormatMobileStar(mobile string) string {
	if len(mobile) <= 7 {
		return mobile[:1] + "****" + mobile[len(mobile)-1:]
	}
	return mobile[:3] + "****" + mobile[len(mobile)-4:]
}

