package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

//RetCodeFromMgmtAPI is return code type
type RetCodeFromMgmtAPI uint32

//Request info
type Request struct {
	BaseURL    string
	pathPrefix string
}

//NewRequest get a Request instance
func NewRequest(server string) *Request {
	pathPrefix := "/"
	if !strings.HasPrefix(server, "http://") {
		server = fmt.Sprintf("http://%s", server)
	}
	return &Request{pathPrefix: pathPrefix,
	BaseURL:server}
}

//SendRequest send request
func (request *Request) SendRequest(method, path string, b []byte) (body []byte) {
	client := &http.Client{}
	path = request.pathPrefix + path

	urlString := fmt.Sprintf("%s%s\n", request.BaseURL, path)

	var req *http.Request
	var err error
	if b == nil {
		req, err = http.NewRequest(method, urlString, nil)
	} else {
		req, err = http.NewRequest(method, urlString, strings.NewReader(string(b)))
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	resp, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") || strings.Contains(err.Error(), "i/o timeout") {
			msg := fmt.Sprintf("Can't connect to api server in %s, is it running?", request.BaseURL)
			PrintAndExit(msg, FailedTemplate)
		} else {
			PrintAndExit(err.Error(), FailedTemplate)
		}
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		PrintAndExit(err.Error(), FailedTemplate)
	}
	CheckStatusCode(resp.StatusCode)
	return body
}
