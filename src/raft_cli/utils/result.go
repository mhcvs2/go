package utils

import (
	"encoding/json"
	"fmt"
	MyLogger "github.com/dbelt/dbcli/logger"
	"strconv"
	"strings"
)

//Result is format api will return
type Result struct {
	RetCode RetCodeFromMgmtAPI `json:"retcode"` //unint32
	Data    interface{}        `json:"data"`
}

//CheckParse if conditionFun return true, program will exit after print origin body we fail to parse.
func CheckParse(body []byte, conditionFun func() bool) {
	if conditionFun() {
		if len(body) > 0 {
			MyLogger.Log.Warn("Fail to parse:")
			PrintAndExit(string(body), FailedTemplate)
		} else {
			PrintAndExit("Empty response body", FailedTemplate)
		}
	}
}

//Parse parse Byte data "body" to struct "res", if failed, print and exit
func Parse(body []byte, res interface{}) {
	err := json.Unmarshal(body, &res)
	CheckParse(body, func() bool {
		if err != nil {
			return true
		}
		return false
	})
}

//CheckResult check data whether is nil, and if RetCode is not equal 0, print reason and exit.
func CheckResult(body []byte) {
	res := Result{}
	Parse(body, &res)
	CheckParse(body, func() bool {
		if res.Data == nil {
			return true
		}
		return false
	})
	if res.RetCode != 0 {
		message := fmt.Sprintf("RetCode is %d, Reason: %s", res.RetCode, strings.TrimSpace(res.Data.(string)))
		PrintAndExit(message, FailedTemplate)
	}
}

//CheckStatusCode check status code 200, if not, print error and exit.
func CheckStatusCode(statusCode int) {
	switch statusCode {
	case 403:
		PrintAndExit("Validation failure.", FailedTemplate)
	default:
		if statusCode >= 300 || statusCode < 200 {
			message := fmt.Sprintf("%s error", strconv.Itoa(statusCode))
			PrintAndExit(message, FailedTemplate)
		}
	}
}

//GetResultData parse result and return result.Data in Bytes
func GetResultData(body []byte) []byte {
	res := Result{}
	Parse(body, &res)
	byteData, _ := json.Marshal(res.Data)
	return byteData
}

//ResultString is a string type Data Result
type ResultString struct {
	RetCode RetCodeFromMgmtAPI `json:"retcode"` //unint32
	Data    string             `json:"data"`
}

//GetResultStringData get data in string type
func GetResultStringData(body []byte) string {
	res := ResultString{}
	Parse(body, &res)
	CheckParse(body, func() bool {
		if res.Data == "" {
			return true
		}
		return false
	})
	return res.Data
}

//ResultStrings is to parse json
type ResultStrings struct {
	RetCode RetCodeFromMgmtAPI `json:"retcode"` //unint32
	Data    []string           `json:"data"`
}

//GetResultStringDatas get result data in slice
func GetResultStringDatas(body []byte) []string {
	res := ResultStrings{}
	Parse(body, &res)
	CheckParse(body, func() bool {
		if len(res.Data) == 0 {
			return true
		}
		return false
	})
	return res.Data
}

//ResultMap is to parse json
type ResultMap struct {
	RetCode RetCodeFromMgmtAPI `json:"retcode"` //unint32
	Data    map[string]string  `json:"data"`
}
