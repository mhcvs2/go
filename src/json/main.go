package main

import (
	"fmt"
	"encoding/json"
	"reflect"
	"github.com/olekukonko/tablewriter"
	"os"
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
	Item1 string
	Item2 int
}

// Implement the String interface for pretty printing Items
func (i Item) String() string {
	return fmt.Sprintf("Item: %s, %d", i.Item1, i.Item2)
}


func Switch(value reflect.Value) {
	kind := value.Kind()
	switch kind {
	case reflect.String:
		fmt.Printf("string: %s\n", value.String())
	case reflect.Struct:
		parseStruct(value)
	case reflect.Slice:
		parseSlice(value)
	case reflect.Interface:
		v := reflect.ValueOf(value.Interface())
		Switch(v)
	case reflect.Map:
		v := reflect.ValueOf(value.Interface())
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			fmt.Println(key.Interface(), strct.Interface())
			Switch(strct)
		}

	default:
		fmt.Printf("unknown type: %s\n", kind.String())
	}
}
func parseSlice(value reflect.Value) {
	size := value.Len()
	for j := 0; j < size; j++ {
		field := value.Index(j)
		Switch(field)
	}
}
func parseStruct(value reflect.Value) {
	size := value.NumField()
	for i := 0; i < size; i++ {
		field := value.Field(i)
		Switch(field)
	}
}

func getShow(x, y int) [][]string {
	res := make([][]string, x)
	for i:=0;i<x;i++{
		res[i] = make([]string, y)
	}
	return res
}

func main() {

	var objs map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &objs); err != nil {
		fmt.Println(err)
		return
	}
	var image map[string]interface{}
	image = objs["images"].(map[string]interface{})
	var runnings map[string]interface{}
	runnings = objs["runnings"].(map[string]interface{})

	show := getShow(5+len(image)+len(runnings),3)
	show[0][0] = "id"
	show[0][1] = objs["id"].(string)
	show[1][0] = "name"
	show[1][1] = objs["name"].(string)
	show[2][0] = "endpoint"
	show[2][1] = objs["endpoint"].(string)
	show[3][0] = "port"
	show[3][1] = fmt.Sprintf("'"+objs["port"].(string)+"'")
	show[4][0] = "state"
	show[4][1] = objs["state"].(string)

	i := 5
	for k, v := range image {
		show[i][0] = "images"
		show[i][1] = k
		show[i][2] = v.(string)
		i++
	}

	for k, v := range runnings {
		show[i][0] = "runnings"
		show[i][1] = k
		show[i][2] = v.(string)
		i++
	}


	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(show)
	table.Render()


}


