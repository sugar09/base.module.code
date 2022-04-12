package base

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

// init 初始化信息
func InitLogs() {
	//异步执行， 1000个缓冲
	logs.Async(1e3)
	logs.SetLogFuncCallDepth(4)
	logs.EnableFuncCallDepth(true)

	//初始化配置信息
	conf := initConf()
	//获取日志类型
	adapterType, _ := beego.AppConfig.String("LogsAdapterType")
	switch adapterType {
	case logs.AdapterConsole:
	case logs.AdapterFile:
	case logs.AdapterMultiFile:
	case logs.AdapterMail:
	case logs.AdapterConn:
	case logs.AdapterEs:
	case logs.AdapterJianLiao:
	case logs.AdapterSlack:
	case logs.AdapterAliLS:
	default:
		adapterType = logs.AdapterMultiFile
	}
	//设置配置信息
	logs.SetLogger(adapterType, conf)
}

type BeegoLogsConf struct {
	//保存文件名
	LogsFileName string `json:"filename,omitempty"`
	//每个文件保存的最大行数
	LogsMaxLines int `json:"maxlines,omitempty"`
	//每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	LogsMaxSize int `json:"maxsize,omitempty"`
	//是否按照每天 logrotate，默认是 true
	LogsDaily interface{} `json:"daily,omitempty"`
	//文件最多保存多少天，默认保存 7 天
	LogsMaxDays int `json:"maxdays,omitempty"`
	//是否开启 logrotate，默认是 true
	LogsRotate interface{} `json:"rotate,omitempty"`
	//日志保存的时候的级别，默认是 Trace 级别
	LogsLevel int `json:"level,omitempty"`
	//日志文件权限
	LogsPerm int `json:"perm,omitempty"`
	//需要单独写入文件的日志级别,设置后命名类似 test.error.log
	LogsSeparate []string `json:"separate,omitempty"`
	//是否开启打印日志彩色打印(需环境支持彩色输出)
	LogsColor interface{} `json:"color,omitempty"`
	//
	//是否每次链接都重新打开链接，默认是 false
	LogsReconnectOnMsg interface{} `json:"reconnectOnMsg,omitempty"`
	//是否自动重新链接地址
	LogsReconnect interface{} `json:"reconnect,omitempty"`
	//网络链接的方式
	LogsNet string `json:"net,omitempty"`
	//网络链接的地址
	LogsAddr string `json:"addr,omitempty"`
	//
	//验证的用户名
	LogsUsername string `json:"username,omitempty"`
	//验证密码
	LogsPassword string `json:"password,omitempty"`
	//发送的邮箱地址
	LogsHost string `json:"host,omitempty"`
	// 邮件需要发送的人，支持多个
	LogsSendTos string `json:"sendTos,omitempty"`
	//发送邮件的标题，默认是 Diagnostic message from server
	LogsSubject string `json:"subject,omitempty"`
	//
	//DNS
	LogsDns string `json:"dns,omitempty"`
}

type BeegoLogs struct {
}

// initConf 初始化配置信息
func initConf() string {
	//创建配置
	conf := &BeegoLogsConf{}

	var data interface{}

	//文件名称
	data, _ = beego.AppConfig.String("LogsFileName")
	if InterfaceToString(data) != "" {
		conf.LogsFileName = InterfaceToString(data)
	}

	//文件保存的最大行数
	data, _ = beego.AppConfig.String("LogsMaxLines")
	if InterfaceToString(data) != "" {
		conf.LogsMaxLines = InterfaceToInt(data)
	}

	//文件保存的最大尺寸
	data, _ = beego.AppConfig.Int("LogsMaxSize")
	if InterfaceToInt(data) > 0 {
		conf.LogsMaxSize = InterfaceToInt(data)
	}

	//是否按照每天
	data, _ = beego.AppConfig.String("LogsDaily")
	if InterfaceToString(data) == "0" {
		conf.LogsDaily = false
	}

	//文件最多保存多少天
	data, _ = beego.AppConfig.Int("LogsMaxDays")
	if InterfaceToInt(data) > 0 {
		conf.LogsMaxDays = InterfaceToInt(data)
	}

	//是否开启 logrotate
	data, _ = beego.AppConfig.String("LogsRotate")
	if InterfaceToString(data) == "0" {
		conf.LogsRotate = false
	}

	//日志保存的时候的级别
	data, _ = beego.AppConfig.Int("LogsLevel")
	if InterfaceToInt(data) > 0 {
		conf.LogsLevel = InterfaceToInt(data)
	}

	//是否开启打印日志彩色打印
	data, _ = beego.AppConfig.String("LogsColor")
	if InterfaceToString(data) == "0" {
		conf.LogsColor = false
	}


	//日志文件权限
	data, _ = beego.AppConfig.String("LogsPerm")
	if InterfaceToString(data) != "" {
		conf.LogsPerm = InterfaceToInt(data)
	}


	//需要单独写入文件的日志级别,设置后命名类似 test.error.log
	data, _ = beego.AppConfig.String("LogsSeparate")
	if InterfaceToString(data) != "" {
		conf.LogsSeparate = strings.Split(InterfaceToString(data), ",")
	}

	//是否每次链接都重新打开链接，默认是 false
	data, _ = beego.AppConfig.String("LogsReconnectOnMsg")
	if InterfaceToString(data) == "1" {
		conf.LogsReconnectOnMsg = true
	}

	//是否每次链接都重新打开链接，默认是 false
	data, _ = beego.AppConfig.String("LogsReconnect")
	if InterfaceToString(data) == "1" {
		conf.LogsReconnect = true
	}

	//发开网络链接的方式，可以使用 tcp、unix、udp 等
	data, _ = beego.AppConfig.String("LogsReconnect")
	if InterfaceToString(data) != "" {
		conf.LogsNet = InterfaceToString(data)
	}

	//网络链接的地址
	data, _ = beego.AppConfig.String("LogsAddr")
	if InterfaceToString(data) != "" {
		conf.LogsAddr = InterfaceToString(data)
	}

	//验证的用户名
	data, _ = beego.AppConfig.String("LogsUsername")
	if InterfaceToString(data) != "" {
		conf.LogsUsername = InterfaceToString(data)
	}

	//验证密码
	data, _ = beego.AppConfig.String("LogsPassword")
	if InterfaceToString(data) != "" {
		conf.LogsPassword = InterfaceToString(data)
	}

	//发送的邮箱地址
	data, _ = beego.AppConfig.String("LogsHost")
	if InterfaceToString(data) != "" {
		conf.LogsHost = InterfaceToString(data)
	}

	//邮件需要发送的人，支持多个
	data, _ = beego.AppConfig.String("LogsSendTos")
	if InterfaceToString(data) != "" {
		conf.LogsSendTos = InterfaceToString(data)
	}

	//发送邮件的标题，默认是 Diagnostic message from server
	data, _ = beego.AppConfig.String("LogsSubject")
	if InterfaceToString(data) != "" {
		conf.LogsSubject = InterfaceToString(data)
	}

	//DNS
	data, _ = beego.AppConfig.String("LogsDns")
	if InterfaceToString(data) != "" {
		conf.LogsDns = InterfaceToString(data)
	}

	//转为json
	bty, _ := json.Marshal(conf)
	return string(bty)
}

// Debug 调式日志
func Debug(v ...interface{}) {
	logs.Debug(v)
}

// Info 普通日志
func Info(v ...interface{}) {
	logs.Info(v)
}

// Error 错误日志
func Error(v ...interface{}) {
	logs.Error(v)
}
