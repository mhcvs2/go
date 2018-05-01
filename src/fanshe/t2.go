package main

import (
	"reflect"
	"runtime"
	"strings"
	"fmt"
)


type MethodMetadata struct {
	Method reflect.Method
	File   string
	Line   int
}


func getFuncMetadata(fn interface{}) (metadata MethodMetadata, err error) {
	v := reflect.ValueOf(fn)

	pc := runtime.FuncForPC(v.Pointer())

	name := strings.TrimRight(pc.Name(), "-fm")
	name = strings.Replace(name, "*", "", 1)

	m := reflect.Method{
		Name:    name,
		PkgPath: "",
		Type:    reflect.TypeOf(fn),
		Func:    v,
		Index:   -1,
	}

	metadata.File, metadata.Line = pc.FileLine(v.Pointer())
	metadata.Method = m

	return
}


type Proxy struct {
	funcs  map[string]MethodMetadata
}

func NewProxy() *Proxy {
	return &Proxy{
		funcs:  make(map[string]MethodMetadata),
	}
}

func (p *Proxy) Method(fn interface{}) (method interface{}) {

	methodName := ""
	methodMetadata, err := getFuncMetadata(fn)
	p.registryFunc(methodMetadata)
	if err != nil {
		panic(err)
	} else {
		methodName = methodMetadata.Method.Name
	}
	if metadata, exist := p.funcs[methodName]; exist {
		method = metadata.Method.Func.Interface()
		return
	}
	return
}

func (p *Proxy) registryFunc(metadata MethodMetadata) {
	fmt.Println("registry---")
	p.funcs[metadata.Method.Name] = metadata
}

//---------------------------------------------------------------------
type Auth struct {}

func (p *Auth) Login(userName, password string) bool {
	fmt.Println("@Login", userName, password)
	if userName == "mhc" && password == "123" {
		return true
	}
	return false
}

//--------------------------------------------------------------
func t21() {
	ret := NewProxy().Method(new(Auth).Login).(func(string, string) bool)("mhc", "haha")
	fmt.Println(ret)
}

func main() {
	t21()
}