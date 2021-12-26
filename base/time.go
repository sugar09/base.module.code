package base

import "time"

//GetNowTimeStamp 返回当前时间的时间戳（秒） 10位数字
func GetNowTimeStamp() int64 {
	//return time.Now().Unix()//跟下面的是等价的
	return time.Now().Local().Unix()
}

//GetNowMsTimeStamp 返回当前时间的毫秒时间戳 13位数字
func GetNowMsTimeStamp() int64 {
	return time.Now().Local().UnixNano() / 1e6
}

