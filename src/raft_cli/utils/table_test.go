package utils

import (
	"fmt"
	"github.com/bouk/monkey"
	assert2 "github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	str = "test"
	i   = 0
	i32 = int32(0)
	i64 = int64(0)
	b   = true
	ss  = []string{}
)

func TestGetString(t *testing.T) {
	assert := assert2.New(t)
	assert.Equal(str, getString(reflect.ValueOf(str)))
	assert.Equal("0", getString(reflect.ValueOf(i)))
	assert.Equal("0", getString(reflect.ValueOf(i32)))
	assert.Equal("0", getString(reflect.ValueOf(i64)))
	assert.Equal("true", getString(reflect.ValueOf(b)))
	assert.Equal("not parse", getString(reflect.ValueOf(ss)))
}

type showStruct struct {
	Name  string `json:"name"`
	IsMan bool   `json:"man"`
	Tall  int64  `json:"tall"`
}

func getShowDatas() []interface{} {
	showDatas := make([]interface{}, 5)
	var name string
	var man bool
	var tall int64
	for i := 0; i < 5; i++ {
		name = fmt.Sprintf("person-%d", i+1)
		if i%2 == 0 {
			man = true
		} else {
			man = false
		}
		tall = int64(150) + int64(i)
		showDatas[i] = showStruct{Name: name, IsMan: man, Tall: tall}
	}
	return showDatas
}

func ExampleShowInTable() {
	showDatas := getShowDatas()
	ShowInTable(showDatas, "Name", "IsMan", "Tall")
	//Output:
	//+----------+-------+------+
	//|   NAME   | ISMAN | TALL |
	//+----------+-------+------+
	//| person-1 | true  |  150 |
	//| person-2 | false |  151 |
	//| person-3 | true  |  152 |
	//| person-4 | false |  153 |
	//| person-5 | true  |  154 |
	//+----------+-------+------+
}

func ExampleShowInTable2() {
	showDatas := getShowDatas()
	ShowInTable(showDatas, "Name", "IsMan")
	//Output:
	//+----------+-------+
	//|   NAME   | ISMAN |
	//+----------+-------+
	//| person-1 | true  |
	//| person-2 | false |
	//| person-3 | true  |
	//| person-4 | false |
	//| person-5 | true  |
	//+----------+-------+
}

func ExampleShowInTable3() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := []interface{}{}
	showDatas = append(showDatas, "test")
	showDatas = append(showDatas, "test2")
	showDatas = append(showDatas, "test3")
	ShowInTable(showDatas, "Name")
	//Output:
	//+-------+
	//| NAME  |
	//+-------+
	//| test  |
	//| test2 |
	//| test3 |
	//+-------+
}

func ExampleChangeSliceType() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := []string{}
	showDatas = append(showDatas, "test")
	showDatas = append(showDatas, "test2")
	showDatas = append(showDatas, "test3")
	showDatas2 := ChangeSliceType(showDatas)
	ShowInTable(showDatas2, "Name")
	//Output:
	//+-------+
	//| NAME  |
	//+-------+
	//| test  |
	//| test2 |
	//| test3 |
	//+-------+
}

func ExampleChangeSliceType2() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := []int{}
	showDatas = append(showDatas, 1)
	showDatas2 := ChangeSliceType(showDatas)
	ShowInTable(showDatas2, "number")
	//Output:
	//+--------+
	//| NUMBER |
	//+--------+
	//|      1 |
	//+--------+
}

func ExampleShowMapInTable() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[interface{}]interface{})
	showDatas["key1"] = "value1"
	showDatas["key2"] = "value2"
	showDatas["key3"] = "value3"
	ShowMapInTable(showDatas,)
	//Output:
	//+------+--------+
	//| NAME | VALUE  |
	//+------+--------+
	//| key1 | value1 |
	//| key2 | value2 |
	//| key3 | value3 |
	//+------+--------+
}

func ExampleShowMapInTable2() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[interface{}]interface{})
	showDatas[1] = "value1"
	showDatas[2] = "value2"
	showDatas[3] = "value3"
	ShowMapInTable(showDatas,"number", "name")
	//Output:
	//+--------+--------+
	//| NUMBER |  NAME  |
	//+--------+--------+
	//|      1 | value1 |
	//|      2 | value2 |
	//|      3 | value3 |
	//+--------+--------+
}

func ExampleShowMapInTable3() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[interface{}]interface{})
	showDatas[1] = "value1"
	ShowMapInTable(showDatas,"number", "name")
	//Output:
	//+--------+--------+
	//| NUMBER |  NAME  |
	//+--------+--------+
	//|      1 | value1 |
	//+--------+--------+
}

func ExampleShowMapInTable4() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[interface{}][]string)
	showDatas[1] = []string{"v1","v2"}
	ShowMapInTable(ChangeMapType(showDatas),"number", "name")
	//Output:
	//+--------+------+
	//| NUMBER | NAME |
	//+--------+------+
	//|      1 | v1   |
	//|        | v2   |
	//+--------+------+
}

func ExampleShowMapInTable5() {
	defer func() {
		recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[interface{}]interface{})
	showDatas["key1"] = "value1"
	showDatas["key2"] = "value2"
	showDatas["key3"] = "value3"
	ShowMapInTable(showDatas,"", "")
	//Output:
	//+------+--------+
	//| key1 | value1 |
	//+------+--------+
	//| key2 | value2 |
	//+------+--------+
	//| key3 | value3 |
	//+------+--------+
}

func ExampleChangeMapType() {
	defer func() {
		//recover()
		monkey.UnpatchAll()
	}()
	Patch()
	showDatas := make(map[int]string)
	showDatas[1] = "value1"
	showDatas2 := ChangeMapType(showDatas)
	ShowMapInTable(showDatas2,"number", "name")
	//Output:
	//+--------+--------+
	//| NUMBER |  NAME  |
	//+--------+--------+
	//|      1 | value1 |
	//+--------+--------+
}
