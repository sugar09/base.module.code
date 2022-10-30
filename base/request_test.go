/**
  @author: sugar
  @since: 2022/10/30
  @desc: //TODO
**/
package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {

	req, err := http.NewRequest("POST", "http://192.168.126.1:8918/v1/login", nil)

	if err != nil {

		fmt.Println(err)
		return
	}

	data := map[string]interface{}{
		"account":  "sugar",
		"password": "e10adc3949ba59abbe56e057f20f883e",
	}

	dataJson, _ := json.Marshal(data)

	fmt.Println(string(dataJson))

	buf := bytes.NewBuffer(dataJson)

	//req.Body.Read(dataJson)

	req.Write(buf)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Post("http://192.168.126.1:8918/v1/login", "application/json", buf)
	//res, err := client.Do(req)

	if err != nil {
		fmt.Println("+++++++++++++", err)
		return
	}

	defer res.Body.Close()

	resb, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(resb), err)

}

func TestRequestPost(t *testing.T) {

	client := &http.Client{}

	data := map[string]interface{}{
		"account":  "sugar",
		"password": "e10adc3949ba59abbe56e057f20f883e",
	}

	dataJson, _ := json.Marshal(data)

	buf := bytes.NewBuffer(dataJson)

	req, err := http.NewRequest("POST", "http://192.168.126.1:8918/v1/login", nil)

	if err != nil {
		fmt.Println("NewRequest:ERR", err)
		return
	}

	//req.Write(buf)

	rc, ok := (io.Reader(buf)).(io.ReadCloser)

	if !ok && buf != nil {
		rc = io.NopCloser(buf)
	}
	req.Body = rc

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("+++++++++++++", err)
		return
	}

	defer res.Body.Close()

	resb, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(resb), err)

}

func TestPostJson(t *testing.T) {

	client := &HttpClient{
		Url: "http://192.168.126.1:8918/v1/loginX",
		RequestData: map[string]interface{}{
			"account":  "sugar",
			"password": "e10adc3949ba59abbe56e057f20f883e",
		},
	}

	res := client.PostJson()

	fmt.Println(res)

	fmt.Println(res.Status)
	fmt.Println(res.Error)
	fmt.Println(string(res.Body))

}
