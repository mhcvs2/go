package utils

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func getString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Int32:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	}
	return "not parse"
}

//ChangeSliceType change slice to type []interface{}, map to map[interface{}]interface{}
func ChangeSliceType(l interface{}) (res []interface{}) {
	value := reflect.ValueOf(l)
	if value.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < value.Len(); i++ {
		res = append(res, value.Index(i).Interface())
	}
	return
}

//ChangeMapType change slice to type []interface{}, map to map[interface{}]interface{}
func ChangeMapType(l interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	value := reflect.ValueOf(l)
	if value.Kind() != reflect.Map {
		return res
	}
	for _, key := range value.MapKeys() {
		res[key.Interface()] = value.MapIndex(key).Interface()
	}
	return res
}

//ShowInTable print []struct/string as table
func ShowInTable(obs []interface{}, keys ...string) {
	var value reflect.Value
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)
	for _, ob := range obs {
		showData := make([]string, len(keys))
		value = reflect.ValueOf(ob)
		switch value.Kind() {
		case reflect.Struct:
			for i, key := range keys {
				field := value.FieldByName(key)
				showData[i] = getString(field)
			}
		default:
			showData[0] = getString(value)
		}
		table.Append(showData)
	}
	table.Render()
}

//ShowMapInTable print map as table
func ShowMapInTable(ob map[interface{}]interface{}, keys ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	var showKeys []string
	if len(keys) == 0 {
		showKeys = []string{"name", "value"}
	} else if keys[0] == "" {
		showKeys = nil
	} else {
		showKeys = keys
	}
	if showKeys != nil {
		table.SetHeader(showKeys)
	} else {
		table.SetRowLine(true)
	}
	for k, v := range ob {
		key := reflect.ValueOf(k)
		value := reflect.ValueOf(v)
		switch value.Kind() {
		case reflect.Slice:
			table.SetRowLine(true)
			keyString := getString(key)
			for i := 0; i < value.Len(); i++ {
				showData := make([]string, 2)
				showData[0] = keyString
				showData[1] = getString(value.Index(i))
				table.Append(showData)
			}
		default:
			showData := make([]string, 2)
			showData[0] = getString(reflect.ValueOf(k))
			showData[1] = getString(reflect.ValueOf(v))
			table.Append(showData)
		}
	}
	table.SetAutoMergeCells(true)
	table.Render()
}

//ShowSplitSliceInTable show split slice in table.
func ShowSplitSliceInTable(ob []interface{}, separator string, keys ...string) {
	table := tablewriter.NewWriter(os.Stdout)
	var showKeys []string
	if len(keys) == 0 {
		showKeys = []string{"name", "value"}
	} else {
		showKeys = keys
	}
	table.SetHeader(showKeys)
	for _, v := range ob {
		showData := make([]string, 2)
		kv := strings.Split(v.(string), separator)
		if len(kv) == 2 {
			showData[0] = strings.TrimSpace(kv[0])
			showData[1] = strings.TrimSpace(kv[1])
		}
		table.Append(showData)
	}
	table.Render()
}
