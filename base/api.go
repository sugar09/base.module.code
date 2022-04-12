package base


import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
)

type Api struct {
	beego.Controller
	//请求的Header参数
	HttpRequestHeaderData map[string]interface{}
	//请求的Get参数
	HttpRequestGetData map[string]interface{}
	//请求的form表单参数
	HttpRequestPostFormData map[string]interface{}
	//请求的JSON参数
	HttpRequestPostJsonData map[string]interface{}
}

// initRequestData 初始化请求信息
func (a *Api) InitRequestData() {
	a.initHttpRequestHeaderData()
	a.initHttpRequestGetData()
	a.initHttpRequestPostFormData()
	a.initHttpRequestPostJsonData()
}

// initHttpRequestPostJsonData 初始化POST Json参数
func (a *Api) initHttpRequestPostJsonData() {
	//可以不用这样写 在配置copyrequestbody = true； a.Ctx.Input.RequestBody
	tmp, _ := ioutil.ReadAll(a.Ctx.Request.Body)
	if len(tmp) > 0 {
		a.HttpRequestPostJsonData = make(map[string]interface{})
		//解析JSON
		e := json.Unmarshal(tmp, &a.HttpRequestPostJsonData)
		if e != nil {
			Error("initHttpRequestPostJsonData:json 参数解析失败，ERR:", e)
		}
	}
}

// initHttpRequestPostFormData 初始化POST form参数
func (a *Api) initHttpRequestPostFormData() {
	tmp := a.Ctx.Request.PostForm
	if len(tmp) > 0 {
		a.HttpRequestPostFormData = make(map[string]interface{})
		for k, v := range tmp {
			a.HttpRequestPostFormData[k] = v[0]
		}
	}
}

// initHttpRequestGetData 初始化Get参数
func (a *Api) initHttpRequestGetData() {
	tmp := a.Ctx.Request.URL.Query()
	if len(tmp) > 0 {
		a.HttpRequestGetData = make(map[string]interface{})
		for k, v := range tmp {
			a.HttpRequestGetData[k] = v[0]
		}
	}
}

// initHttpRequestGetData 初始化Get参数
func (a *Api) initHttpRequestHeaderData() {
	tmp := a.Ctx.Request.Header
	if len(tmp) > 0 {
		a.HttpRequestHeaderData = make(map[string]interface{})
		for k, v := range tmp {
			a.HttpRequestHeaderData[k] = v[0]
		}
	}
}
