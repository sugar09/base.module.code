package logs

type LogInterface interface {
	//Debug 调试日志
	Debug(...interface{})

	//Info 普通日志
	Info(...interface{})

	//Error 错误日志
	Error(...interface{})

}
