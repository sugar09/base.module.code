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

//GetNowUsTimeStamp 返回当前时间的微秒时间戳 16位数字
func GetNowUsTimeStamp() int64 {
	return time.Now().Local().UnixNano() / 1e3
}

//GetNowNsTimeStamp 返回当前时间的纳秒时间戳 19位数字
func GetNowNsTimeStamp() int64 {
	return time.Now().Local().UnixNano()
}