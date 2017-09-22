package main

import (
	"fmt"
	"encoding/json"
	"reflect"

)


//var jsonBytes = []byte(`[
//{
//    "key1":{
//        "Item1": "Value1",
//        "Item2": 1},
//    "key2":{
//        "Item1": "Value2",
//        "Item2": 2},
//    "key3":{
//        "Item1": "Value3",
//        "Item2": 3},
//    "key4":["test1","test2"]
//}]`)

var jsonBytes = []byte(`
{
    "endpoint": "redis-testredis.svc.41d254d6-6b12-43a6-d8b8-d79b0662346b.bst-1.cns.bstjpc.com",
    "id": "testredis",
    "images": {
        "dbproxy": "a7b8d9a51803",
        "redis": "5962c9dddbe0"
    },
    "name": "testrediscluster",
    "port": "6379",
    "runnings": {
        "dbproxy": "1/1:[6e29bccdd5cc]",
        "mgmtapi": "1/1:[54ca4fb07d91]",
        "redis": "3/3:[2ccbf521bc99,794c61b820dd,e81d33e43533]"
    },
    "state": "avaiable"
}`)

// Item struct; we want to create these from the JSON above
type Item struct {
	Str string
	Map map[string]Item
}

func Switch(objs interface{}, resMap map[string]Item) {
	switch expr {

	}
	//value := reflect.ValueOf(objs)
	//if value.Kind() == reflect.Map {
	//	v := reflect.ValueOf(value.Interface())
	//	for _, key := range v.MapKeys() {
	//		strct := v.MapIndex(key)
	//		switch strct.Kind() {
	//		case reflect.String:
	//			strct.(string)
	//		}
	//
	//	}
	//} else {
	//	fmt.Printf("unknown type: %s\n", value.String())
	//}
}

func main() {

	var objs interface{}
	if err := json.Unmarshal(jsonBytes, &objs); err != nil {
		fmt.Println(err)
		return
	}
	resMap := Item{}
	Switch(objs, &resMap)

}


