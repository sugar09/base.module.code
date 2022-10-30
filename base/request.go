/**
  @author: sugar
  @since: 2022/10/30
  @desc: //TODO
**/
package base

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//
//  HttpClient
//  @Description: Http请求客户端
//
type HttpClient struct {
	Url           string                                             //请求地址
	Header        map[string]string                                  //请求头
	body          io.Reader                                          //请求body
	RequestData   map[string]interface{}                             //请求数据
	CheckRedirect func(req *http.Request, via []*http.Request) error //检查重定向
	Transport     *http.Transport                                    //Transport
}

//
//  HttpResponse
//  @Description: http请求返回信息
//
type HttpResponse struct {
	Status int    //http 状态码
	Body   []byte // 返回body
	Error  error  // 错误信息
}

//
//  PostJson
//  @Description: Post json请求
//  @receiver that 当前结构体
//  @return res *HttpResponse 返回信息
//
func (that *HttpClient) PostJson() (res *HttpResponse) {
	var err error
	res = &HttpResponse{}
	that.getBody()
	req, err := http.NewRequest("POST", that.Url, that.body)
	if err != nil {
		res.Error = err
		return
	}
	//默认
	req.Header.Set("Content-Type", "application/json")

	that.setHeader(req)

	resp, err := that.request(req)

	res.Status = resp.StatusCode

	if err != nil {
		res.Error = err
	}
	defer resp.Body.Close()
	res.Body, err = ioutil.ReadAll(resp.Body)
	return
}

//
//  Get
//  @Description: Get请求
//  @receiver that 当前结构体
//  @return res *HttpResponse 返回信息
//
func (that *HttpClient) Get() (res *HttpResponse) {
	var err error
	res = &HttpResponse{}
	req, err := http.NewRequest("GET", that.Url, nil)
	if err != nil {
		res.Error = err
		return
	}
	that.setHeader(req)

	resp, err := that.request(req)

	res.Status = resp.StatusCode

	if err != nil {
		res.Error = err
	}
	defer resp.Body.Close()

	res.Body, err = ioutil.ReadAll(resp.Body)
	return
}

//
//  getBody
//  @Description: 获取请求信息
//  @receiver that 当前结构体
//
func (that *HttpClient) getBody() {
	if that.RequestData != nil {

		dataJson, err := json.Marshal(that.RequestData)

		if err == nil {
			that.body = bytes.NewBuffer(dataJson)
		}
	}
}

//
//  setHeader
//  @Description: 设置头部信息
//  @receiver that 当前结构体
//  @param req *http.Request 请求体
//
func (that HttpClient) setHeader(req *http.Request) {
	if that.Header != nil {
		for k, v := range that.Header {
			req.Header.Set(k, v)
		}
	}
}

//
//  request
//  @Description: 请求
//  @receiver that 当前结构体
//  @param req *http.Request 请求体
//  @return *http.Response 请求返回信息
//  @return error 错误信息
//
func (that HttpClient) request(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	if that.CheckRedirect != nil {
		client.CheckRedirect = that.CheckRedirect
	}

	if that.Transport != nil {
		client.Transport = that.Transport
	}

	return client.Do(req)
}
