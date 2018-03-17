package main

import (
	"os"
	"text/template"
	"fmt"
	"bytes"
)

func t11() {
	name := "waynehu"
	tmpl, err := template.New("test").Parse("hello, {{.}}") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name)  //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}
//hello, waynehu


type Inventory struct {
	Material string
	Count    uint
}

func t12() {
	sweaters := Inventory{"wool", 17}
	muban := "{{.Count}} items are made of {{.Material}}"
	tmpl, err := template.New("test").Parse(muban)  //建立一个模板
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters) //将struct与模板合成，合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}

//17 items are made of wool

func t13() {
	////一个模板可以有多种，以Name来区分
	//muban_eng := "{{.Count}} items are made of {{.Material}}"
	//muban_chn := "{{.Material}}做了{{.Count}}个项目"
	////建立一个模板的名称是china，模板的内容是muban_chn字符串
	//tmpl := template.New("china")
	//tmpl, _ = tmpl.Parse(muban_chn)
	////建立一个模板的名称是english，模板的内容是muban_eng字符串
	//tmpl = tmpl.New("english")
	//tmpl, _ = tmpl.Parse(muban_eng)
	////将struct与模板合成，用名字是china的模板进行合成，结果放到os.Stdout里，内容为“wool做了17个项目”
	//tmpl.ExecuteTemplate(os.Stdout, "china", sweaters)
	////将struct与模板合成，用名字是china的模板进行合成，结果放到os.Stdout里，内容为“17 items are made of wool”
	//err = tmpl.ExecuteTemplate(os.Stdout, "english", sweaters)
	//
	//
	//tmpl, err = template.New("english")
	//fmt.Println(tmpl.Name())  //打印出english
	//tmpl, err = tmpl.New("china")
	//fmt.Println(tmpl.Name())  //打印出china
	//tmpl=tmpl.Lookup("english")//必须要有返回，否则不生效
	//fmt.Println(tmpl.Name())  //打印出english
}

func t14() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.ParseFiles("/root/github/go/src/template/t1.txt")
	var res = &bytes.Buffer{}
	err = tmpl.Execute(res, sweaters)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res.String())
	}
}

func main() {
	t14()
}