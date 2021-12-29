package base
//正则验证
import (
	"fmt"
	"regexp"
)

//CheckInt 验证数字
var CheckInt = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^\d+$`, fmt.Sprint(str))
	return match
}

//CheckInts 验证数字，多个逗号拼接的IDs
var CheckInts = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^\d*(\d+\,)*\d+$`, fmt.Sprint(str))
	return match
}

var CheckEnStr = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z\d]$+`, fmt.Sprint(str))
	return match
}

// CheckWStr 验证\w的字符
var CheckWStr = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^\w+$`, fmt.Sprint(str))
	return match
}

// CheckWStrWithLen 验证\w的指定长度的字符
var CheckWStrWithLen = func(str interface{}, len int) bool {
	match, _ := regexp.MatchString(fmt.Sprintf(`^\w{%d}$`, len), fmt.Sprint(str))
	return match
}

// CheckWStrWithRan 验证\w的指定范围长度的字符
var CheckWStrWithRan = func(str interface{}, min int, max int) bool {
	match, _ := regexp.MatchString(fmt.Sprintf(`^\w{%d,%d}$`, min, max), fmt.Sprint(str))
	return match
}

// CheckUniqueStr 验证\w的字符
var CheckUniqueStr = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^[\w\-]+$`, fmt.Sprint(str))
	return match
}

// CheckUniqueStrWithLen 验证\w的指定长度的字符
var CheckUniqueStrWithLen = func(str interface{}, len int) bool {
	match, _ := regexp.MatchString(fmt.Sprintf(`^[\w\-]{%d}$`, len), fmt.Sprint(str))
	return match
}

// CheckUniqueStrWithRan 验证\w的指定范围长度的字符
var CheckUniqueStrWithRan = func(str interface{}, min int, max int) bool {
	match, _ := regexp.MatchString(fmt.Sprintf(`^[\w\-]{%d,%d}$`, min, max), fmt.Sprint(str))
	return match
}

//CheckCnEnNumStr 验证英文、数字、中文、常用符号字符串
var CheckCnEnNumStr = func(str interface{}) bool {
	//^[\w\u4E00-\u9FA5A\s\,\.\!\?\+\-\(\)\、\，\。\！\？\（\）]+$
	match, _ := regexp.MatchString(`^[\w\p{Han}\s\,\.\!\?\+\-\(\)\\、，。！？（）@#￥%/|]+$`, fmt.Sprint(str))
	return match
}


//CheckEmail 验证电子邮箱
var CheckEmail = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9]+([._\-][a-zA-Z0-9]+)*@([a-zA-Z0-9]+([._\-][a-zA-Z0-9]+))+$`, fmt.Sprint(str))
	return match
}

//CheckPhone 验证手机号码
var CheckPhone = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^1[\d]{10}$`, fmt.Sprint(str))
	return match
}



//CheckTimestamp 验证10位秒时间戳
var CheckTimestamp = func(str interface{}) bool {
	match, _ := regexp.MatchString(`^\d{10}$`, fmt.Sprint(str))
	return match
}

//CheckDateTime 验证时间 格式 2020-04-24 15:00:00
func CheckDateTime(str interface{}) bool {
	match, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}$`, fmt.Sprint(str))
	return match
}