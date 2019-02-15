package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func WebServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func login(w http.ResponseWriter, req *http.Request) {

	fmt.Println("method:", req.Method) //获取请求的方法

	if req.Method == "GET" {

		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))


		t, _ := template.ParseFiles("src/html/login.html")
		log.Println(t.Execute(w, token))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		req.ParseForm()

		token := req.Form.Get("token")
		if token != "" {
			//验证token合法性
		} else {
			//不存在token则报错
		}

		fmt.Println("username length:", len(req.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(req.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(req.Form.Get("password")))
		template.HTMLEscape(w, []byte(req.Form.Get("username"))) //输出到客户端
	}
}

func upload(w http.ResponseWriter, req *http.Request) {

	fmt.Println("method:", req.Method) //获取请求的方法

	if req.Method == "GET" {

		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))


		t, _ := template.ParseFiles("src/html/upload.html")
		log.Println(t.Execute(w, token))

	} else {

		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer  f.Close()
		io.Copy(f, file)
	}
}

func postFile(filename string, targetUrl string) error {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
	return nil
}

func main()  {
	WebServerBase()

	time.Sleep(10)
	targetUrl := "http://localhost:8081/upload"
	filename := "./BAT算法面试题(2).pdf"
	postFile(filename, targetUrl)
}

