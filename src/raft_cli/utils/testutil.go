package utils

import (
	"encoding/json"
	"fmt"
	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//SuccessResponseBytes is SuccessResponseBytes
var (
	appid           = "test"
	appsceret       = "ee1d7336-73e3-452b-9e63-fdeaf2dccde6"
	successResponse = Result{
		RetCode: 0,
		Data:    "success",
	}
	SuccessResponseBytes, _ = json.Marshal(successResponse)
	failedResponse          = Result{
		RetCode: 1,
		Data:    "failed",
	}
	failedResponseBytes, _ = json.Marshal(failedResponse)
)

// Patch exit function
func Patch() {
	monkey.Patch(PrintAndExit, func(message, template string) {
		fmt.Println(message)
	})
}

//InitServer init a httptest server
func InitServer(t *testing.T, path, method string, response []byte, bodyStruct interface{}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert := assert.New(t)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		assert.Equal(method, r.Method)
		assert.Equal(path, r.URL.EscapedPath())
		r.ParseForm()
		formAppid := r.Form.Get("appid")
		if bodyStruct != nil {
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bodyStruct)
		}
		assert.Equal(appid, formAppid)
	}))
	return ts
}

//InitServer2 init a httptest server without check
func InitServer2(path, method string, response []byte, bodyStruct interface{}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		r.ParseForm()
		if bodyStruct != nil {
			body, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bodyStruct)
		}
	}))
	os.Setenv("MGMTSERVER", ts.URL)
	os.Setenv("OPSSERVER", ts.URL)
	os.Setenv("APPID", appid)
	os.Setenv("APPSCERET", appsceret)
	return ts
}
