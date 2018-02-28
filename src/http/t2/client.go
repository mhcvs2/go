package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
)

func httpGet() {
	resp, err := http.Get("http://localhost:3322")
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

func httpPut(message string) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3322", strings.NewReader(message))
	resp, err := client.Do(req)
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

func httpDelete() {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3322", nil)
	resp, err := client.Do(req)
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
	resp, err := http.Post("http://localhost:3322",
		"application/x-www-form-urlencoded",
		strings.NewReader("exit"))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(resp)
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
	//httpPut("hello world")
	//httpGet()
	//httpDelete()
	httpPost()
}