package main

import (
	"fmt"
	"reflect"
)

func t1() {
	var a interface{}
	a = "dsfsdf"
	fmt.Println(a)
	value := reflect.ValueOf(a)
	fmt.Println(value.String())
	fmt.Println(value.Kind())
	fmt.Println(value.Len())
	fmt.Println(value.IsValid())
	fmt.Println(value.Interface())
	fmt.Println(value.Type())
}
/*
dsfsdf
dsfsdf
string
6
true
dsfsdf
string
*/

func t2() {
	var a2 interface{}
	a2 = []string{"a","b","c"}
	value := reflect.ValueOf(a2)
	fmt.Println(value.Kind())
	fmt.Println(value.Len())
	fmt.Println("------------")
	for i:=0;i<value.Len();i++ {
		fmt.Println(value.Index(i))
	}
	fmt.Println("------------")
	fmt.Println(value.Index(0).Kind())
}
/*
slice
3
------------
a
b
c
------------
string
*/

func t3() {
	var a3 interface{}
	b3 := make(map[string]string)
	b3["sdff"] = "ddd"
	b3["aaa"] = "sss"
	a3 = b3

	value := reflect.ValueOf(a3)
	fmt.Println(value.Kind())
	fmt.Println(value.MapKeys())
	fmt.Println("-------------")
	for _, key := range value.MapKeys() {
		fmt.Println(value.MapIndex(key).String())
	}
	fmt.Println("------------")
	value.SetMapIndex(reflect.ValueOf("hahah"),reflect.ValueOf("yaya"))
	fmt.Println(a3)
	fmt.Println(value.CanSet())
}
/*
map
[sdff aaa]
-------------
ddd
sss
------------
map[sdff:ddd aaa:sss hahah:yaya]
false
*/

type Test4 struct {
	X int
	Y string
}

func t4() {
	var a4 interface{}
	a4 = Test4{122, "sadad"}
	value := reflect.ValueOf(a4)

	fmt.Println(value.Kind())
	fmt.Println(value.Type().Name())
	fmt.Println(value.NumField())
	fmt.Println("-------------")
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i))
	}
	fmt.Println("-------------")
	fmt.Println(value.FieldByIndex([]int{0}))
	fmt.Println(value.FieldByName("Xa"))
	fmt.Println(reflect.Value{})
	if value.FieldByName("Xa") == (reflect.Value{}){
		fmt.Println("kong")
	}
}
//
//struct
//Test4
//2
//-------------
//122
//sadad
//-------------
//122
//<invalid reflect.Value>
//<invalid reflect.Value>
//kong


type Test5 struct {
	X int
	Y string
}

func (t Test5) TestMethod() {
	fmt.Println("I am test5 method")
}
func (t Test5) Test2Method() {
	fmt.Println("X:",t.X,"Y:",t.Y)
}
func (t Test5) Test3Method(name string) {
	fmt.Println("hello ", name)
}

func (t *Test5) Test4Method() {
	fmt.Println("hello test4")
}

func t5() {
	var a5 interface{}
	a5 = Test5{111, "sss"}
	value := reflect.ValueOf(a5)
	fmt.Println(value.NumMethod())
	value.Method(0).Call([]reflect.Value{})
	value.MethodByName("TestMethod").Call([]reflect.Value{})
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf("mhc")
	value.MethodByName("Test3Method").Call(args)
	fmt.Println("----------------------------------")
	var a52 *Test5
	a52 = &Test5{11,"ha"}
	value = reflect.ValueOf(a52)
	fmt.Println(value.NumMethod())
	value.Method(0).Call([]reflect.Value{})
	value.MethodByName("TestMethod").Call([]reflect.Value{})
	args = make([]reflect.Value, 1)
	args[0] = reflect.ValueOf("mhc")
	value.MethodByName("Test3Method").Call(args)
	value.MethodByName("Test4Method").Call([]reflect.Value{})
	method := value.MethodByName("sdfsdf")
	if (method == reflect.Value{}){
		fmt.Println("mei")
	}
}
/*
3
X: 111 Y: sss
I am test5 method
hello  mhc
----------------------------------
4
X: 11 Y: ha
I am test5 method
hello  mhc
hello test4
*/

func t6() {
	p := make(map[string]string)
	p["a"] = "aaa"
	value := reflect.ValueOf(p)
	fmt.Println(value.Kind())
	fmt.Println(value.MapKeys())
	for _, key := range value.MapKeys(){
		fmt.Println(value.MapIndex(key))
	}
}

func t7() {
	p := []string{"a","b"}
	value := reflect.ValueOf(p)
	fmt.Println(value.Kind())
	fmt.Println(value.Len())
	fmt.Println(value.Index(0))
}


func main() {
	t4()
}
