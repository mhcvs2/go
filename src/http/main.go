package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
)

func httpGet() {
	resp, err := http.Get("http://localhost:8080/v1/mhc")
	if err != nil {
		fmt.Println("error:", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://localhost:8080/v1/mhc",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println("error:", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(body))
}

func httpFormPost() {
	//这里添加post的body内容
	data := make(url.Values)
	data["user"] = []string{"name"}
	data["pass"] = []string{"aaaaa"}

	//把post表单发送给目标服务器
	res, err := http.PostForm("http://localhost:8081/login", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(body))
}

func main() {
	httpFormPost()
}