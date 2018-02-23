package main

import (
	"net/url"
	"fmt"
)

func t11() {
	//aUrl, err := url.Parse("http://www.baidu.com/api")  // true
	aUrl, err := url.Parse("baidu.com/api")         // flase
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aUrl)
	fmt.Println(aUrl.IsAbs())
}
//baidu.com/api
//false

func t12() {
	aUrl, _ := url.Parse("http://www.baidu.com/api")
	bUrl, _ := url.Parse("v1/haha")
	fmt.Println(aUrl.ResolveReference(bUrl))
}
// http://www.baidu.com/v1/haha

func main() {
	t12()
}
