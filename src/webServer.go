package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main()  {
	webServerBase()
}

func webServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", loginTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func proxyTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("proxyTask is running...")
	proxy := "c3NyOi8vTVRrMExqRTBOeTR6TlM0NU5Ub3lNek16T205eWFXZHBianBoWlhNdE1qVTJMV05tWWpwd2JHRnBianBrUjJScldWZHNjMkZSTHo5dlltWnpjR0Z5WVcwOUpuQnliM1J2Y0dGeVlXMDlKbkpsYldGeWEzTTlaRU0xZEZwVE9UQmFNbEpvWVZkNGNEVnZMVkUxVERaaVNVOVRYMmhQWlRsc0xXRlhjaTFUTkdsbFpUbHJaV1ZpZEU5cFgyNW5KbWR5YjNWd1BRCgpzczovL1lXVnpMVEkxTmkxalptSTZaVWxYTUVSdWF6WTVORFUwWlRadVUzZDFjM0IyT1VSdFV6SXdNWFJSTUVSQU1UY3lMakV3TkM0MU1DNDJOem80TURrMwoKc3M6Ly9jbU0wTFcxa05Ub3lPVEp1WVRGemFFQXhPVGd1TWk0eU5UTXVNemM2TVRBMU5RCgp2bWVzczovL1kyaGhZMmhoTWpBdGNHOXNlVEV6TURVNllUTmpNV015WVRrdFpXTTFPQzAwWWpSakxUZzBPRGt0TXprMk5UUmtZMll6TkRBNFFIWXlMblJuWkdGcGJHbG1jbVZsTG5oNWVqbzBORE09P25ldHdvcms9d3Mmd3NwYXRoPXYyLnRnZGFpbGlmcmVlLnh5eiZhaWQ9MCZ0bHM9MSZhbGxvd0luc2VjdXJlPTEmbXV4PTAmbXV4Q29uY3VycmVuY3k9OCZyZW1hcms9dC5tZS90Z2RhaWxpJUU2JThGJTkwJUU0JUJFJTlCJTIwJUU0JUJGJTg0JUU3JUJEJTk3JUU2JTk2JUFGJUU2JTk3JUEwJUU5JTk5JTkwJUU2JUI1JTgxJUU5JTg3JThGCgp2bWVzczovL1lXVnpMVEV5T0MxalptSTZOMkU1WW1ZNU9HWXRaalZoWmkwMFlqQmlMV0UzT1RNdE16UmhNRFpoTURNNFlUUTRRSE5sY25abGNqRXVhSFZoYm1kamIyNW5MbTFzT2pVMk56Zz0/bmV0d29yaz10Y3AmYWlkPTY0JnRscz0wJmFsbG93SW5zZWN1cmU9MSZtdXg9MCZtdXhDb25jdXJyZW5jeT04JnJlbWFyaz0wJTIwU2lsaWNvbiUyMFZhbGxleSUyMFYycmF5JTIwVENQ"
	fmt.Fprint(w, proxy)
}

func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")

	//模拟延时
	time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	result := NewBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}
