package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

// Request 请求体
type Request struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Method  string
	Body    string
}

// Do 发送请求
func (r *Request) Do() {
	if r.Method == GET {
		// TODO: 实现GET请求
		resp, err := http.Get(buildGetUrl(r.Url, r.Params))
		if err != nil {
			// TODO: 处理错误
		}
		defer resp.Body.Close()
		// 读取响应
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// 打印响应
		fmt.Println(string(body))
	} else {
		// TODO: 实现POST请求
		resp, err := http.Post(r.Url, "application/json", bytes.NewBufferString(r.Body))
		if err != nil {
			// TODO: 处理错误
		}
		defer resp.Body.Close()
		// 读取响应
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		// 打印响应
		fmt.Println(string(body))
	}

}

func buildGetUrl(url string, params map[string]string) string {
	if len(params) == 0 {
		return url
	}
	url += "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}
	return url[:len(url)-1]
}
