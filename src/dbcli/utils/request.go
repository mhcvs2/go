package utils

import (
	"time"
	"net/url"
	"github.com/astaxie/beego/plugins/apiauth"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	MyLogger "dbcli/logger"
	"strconv"
	"os"
)

func PreRun() {
	appid := os.Getenv("APPID")
	appsecret := os.Getenv("APPSCERET")
	if appid == "" || appsecret == "" {
		PrintAndExit("ss", preRunTemplate)
	}
}

type Request struct {
	baseUrl string
	path_prefix string
	appid string
	appsecret string
}

func NewRequest() *Request {
	baseUrl := os.Getenv("MGMTSERVER")
	if baseUrl == "" {
		baseUrl = "http://127.0.0.1:8080"
	}
	path_prefix := "/v1/"
	appid := os.Getenv("APPID")
	//appsecret := "ee1d7336-73e3-452b-9e63-fdeaf2dccde6"
	appsecret := os.Getenv("APPSCERET")
	return &Request{baseUrl:baseUrl, path_prefix:path_prefix, appid:appid, appsecret:appsecret}
}

func (request *Request) SendRequest(method, path string, b []byte) (body []byte) {
	client := &http.Client{}
	path = request.path_prefix + path
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	params := url.Values{"appid": {request.appid}, "timestamp": {timestamp}}
	signature := url.QueryEscape(apiauth.Signature(request.appsecret, method, params, path))

	urlString := fmt.Sprintf("%s%s?%s%s&&%s%s&&%s%s", request.baseUrl, path, "appid=", request.appid, "timestamp=", url.QueryEscape(timestamp), "signature=", signature)
	MyLogger.Log.Debugf("url=", urlString)

	var req *http.Request
	if b == nil {
		req, _ = http.NewRequest(method, urlString, nil)
	} else {
		req, _ = http.NewRequest(method, urlString, strings.NewReader(string(b)))
	}

	resp, err := client.Do(req)
	if err != nil {
		PrintAndExit(err.Error(), FailedTemplate)
	}

	defer resp.Body.Close()
	if resp.StatusCode == 403 {
		PrintAndExit("Permission denied", FailedTemplate)
	}
	if resp.StatusCode != 200 {
		PrintAndExit(strconv.Itoa(resp.StatusCode), FailedTemplate)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintAndExit(err.Error(), FailedTemplate)
	}
	return body
}
