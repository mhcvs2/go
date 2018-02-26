package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequest_SendRequest(t *testing.T) {
	ts := InitServer(t, "/v1/test", "GET", SuccessResponseBytes, nil)
	defer ts.Close()
	request := NewRequest()
	resultBytes := request.SendRequest("GET", "test", []byte{})
	assert.Equal(t, GetResultStringData(resultBytes), "success")
}

type RequestTest struct {
	Name string `json:"name"`
}

func TestRequest_SendRequest2(t *testing.T) {
	req := RequestTest{Name: "mhc"}
	reqByte, _ := json.Marshal(req)
	reqServer := RequestTest{}
	ts := InitServer(t, "/v1/test", "POST", SuccessResponseBytes, &reqServer)
	defer ts.Close()
	request := NewRequest()
	resultBytes := request.SendRequest("POST", "test", reqByte)
	assert.Equal(t, GetResultStringData(resultBytes), "success")
	assert.Equal(t, req.Name, reqServer.Name)
}
