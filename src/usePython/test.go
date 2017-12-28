package main

import (
	"github.com/sbinet/go-python"
	"fmt"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING

func main() {
	// import hello
	InsertBeforeSysPath("/usr/lib/python2.7/site-packages")
	hello := ImportModule("/root/github/python/re_test", "d")

	// print(hello.b)
	t2 := hello.GetAttrString("t21")

	// b(*args)
	res := t2.CallFunction()
	var item *python.PyObject
	for i:=0; i<python.PyList_GET_SIZE(res); i++ {
		item = python.PyList_GET_ITEM(res, i)
		fmt.Println(python.PyString_AS_STRING(item))
	}
}
//Mountain View
//CA
//94040

// InsertBeforeSysPath will add given dir to python import path
func InsertBeforeSysPath(p string) string {
	sysModule := python.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	python.PyList_Insert(path, 0, PyStr(p))
	return GoStr(path.Repr())
}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")                    // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))                     // path.insert(0, dir)
	return python.PyImport_ImportModule(name)            // return __import__(name)
}
