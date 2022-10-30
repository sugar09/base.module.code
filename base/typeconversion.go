package base

//基本类型转换
import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"strconv"
	"strings"
)

// InterfaceToString interface类型转换String
var InterfaceToString = func(val interface{}) string {
	tf := reflect.TypeOf(val)
	if tf == nil {
		return ""
	}
	switch tf.Kind() {
	case reflect.Map:
		fmt.Println("Map:", val)
		return ""
	case reflect.Chan:
		fmt.Println("Map:", val)
		return ""
	}
	fmt.Println("反射类型", reflect.TypeOf(val))
	fmt.Println("反射类型", reflect.TypeOf(val).Kind())
	return fmt.Sprint(val)
}

// InterfaceToInt64 interface类型转int64
var InterfaceToInt64 = func(val interface{}) int64 {
	i, _ := strconv.ParseInt(InterfaceToString(val), 10, 64)
	return i
}

// InterfaceToInt32 interface类型转int32
var InterfaceToInt32 = func(val interface{}) int32 {
	return int32(InterfaceToInt64(val))
}

// InterfaceToInt interface类型转int
var InterfaceToInt = func(val interface{}) int {
	return int(InterfaceToInt64(val))
}

// InterfaceToFloat64 interface类型转float64
var InterfaceToFloat64 = func(val interface{}) float64 {
	f, _ := strconv.ParseFloat(InterfaceToString(val), 64)
	return f
}

// InterfaceToFloat32 interface类型转float32
var InterfaceToFloat32 = func(val interface{}) float32 {
	return float32(InterfaceToFloat64(val))
}

// InterfaceToMapWithStringInterface interface类型转为map
var InterfaceToMapWithStringInterface = func(val interface{}) map[string]interface{} {
	//检测变量类型是否为map
	if reflect.TypeOf(val).Kind() == reflect.Map {
		return val.(map[string]interface{})
	}
	return map[string]interface{}{}
}

var RebuildScienceNumToString = func(numStr string) string {
	if strings.Contains(numStr, "e+") {
		//是否科学计数法
		decimalNum, err := decimal.NewFromString(numStr)
		if err != nil {
			Error("获取科学计数数值出错", err)
			return numStr
		}
		decimalStr := decimalNum.String()
		return decimalStr
	}
	return numStr
}

