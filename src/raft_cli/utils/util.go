package utils

import (
	"io"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dbelt/dbcli/config"
	myLogger "github.com/dbelt/dbcli/logger"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// GetAppName get this app name
func GetAppName() string {
	return config.App.AppName
}

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// FILE returns the file name in which the function was invoked
func FILE() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

//Contains check whether string contain kvs
func Contains(s string, kvs ...string) bool {
	for _, kv := range kvs {
		if !strings.Contains(strings.ToLower(s), strings.ToLower(strings.TrimSpace(kv))) {
			return false
		}
	}
	return true
}

//SliceFilter return a new slice which is filtered by keywords("kw1&&kw2...") and a condition function.
//Target can be a slice of string or struct, when struct, we expect first keywords is target field name of struct
//and expect it's value in string type.
func SliceFilter(target []interface{}, keywords string, conditionFun ...func(s reflect.Value) bool) (res []interface{}) {
	kws := strings.Split(keywords, "&")
	var targetString string
	var targetKws []string
	for _, s := range target {
		value := reflect.ValueOf(s)
		switch value.Kind() {
		case reflect.String:
			targetString = value.String()
			targetKws = kws
		case reflect.Struct:
			if len(kws) < 1 {
				continue
			}
			targetString = value.FieldByName(kws[0]).String()
			targetKws = kws[1:]
		}
		if strings.TrimSpace(targetString) == "" || (len(conditionFun) > 0 && conditionFun[0](value) == true) ||
			(len(targetKws) > 0 && strings.TrimSpace(targetKws[0]) != "" && !Contains(targetString, targetKws...)) {
			continue
		}
		res = append(res, s)
	}
	return
}

//SliceToMap transfer slice to map
func SliceToMap(ss []interface{}, separator string) map[string]string {
	res := make(map[string]string)
	for _, s := range ss {
		kv := strings.Split(s.(string), separator)
		if len(kv) == 2 {
			res[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return res
}

//PrintJSON print json as good format
func PrintJSON(showByte []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, showByte, "", "    ")
	if err != nil {
		myLogger.Log.Error("Invilad json stytle:")
		PrintAndExit(string(showByte), FailedTemplate)
	}
	PrintAndExit(string(out.Bytes()), SuccessTemplate)
}

//CheckNil check the flag whether is nil
func CheckNil(name, value string) bool {
	if value == "" {
		message := fmt.Sprintf("Please spectify %s", "--"+name)
		Tmpl(FailedTemplate, message)
		return true
	}
	return false
}

//CheckNilMulti check the flag whether is nil
func CheckNilMulti(kvs map[string]string) (res bool) {
	for k, v := range kvs {
		if v == "" {
			message := fmt.Sprintf("Please spectify %s", "--"+k)
			Tmpl(FailedTemplate, message)
			res = true
		}
	}
	return
}

//SplitLines split string with '/n'
func SplitLines(data string) []string {
	n := []byte{92, 110}
	var res1 []string
	var res2 []string
	res1 = strings.Split(data, string(n))
	res2 = strings.Split(data, "\n")
	if len(res1) > len(res2) {
		return res1
	}
	return res2
}


//PathExists return file/directory whether exist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}