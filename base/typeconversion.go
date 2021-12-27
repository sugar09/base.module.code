package base
//基本类型转换
import (
	"fmt"
	"strconv"
)

// InterfaceToString interface类型转换String
var InterfaceToString = func(val interface{}) string {
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
	return val.(map[string]interface{})
}
