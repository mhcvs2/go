package utils

import (
	"encoding/json"
	"fmt"
	"github.com/bouk/monkey"
)

func ExampleCheckParse() {
	defer monkey.UnpatchAll()
	Patch()
	body := []byte("not empty")
	CheckParse(body, func() bool {
		return true
	})
	// Output: not empty
}

func ExampleCheckParse2() {
	defer monkey.UnpatchAll()
	Patch()
	body := []byte("")
	CheckParse(body, func() bool {
		return true
	})
	// Output: Empty response body
}

func ExampleParse_Success() {
	defer monkey.UnpatchAll()
	Patch()
	res := Result{}
	Parse(SuccessResponseBytes, &res)
	fmt.Println(res.Data)
	// Output: success
}

func ExampleParse_Failed() {
	defer monkey.UnpatchAll()
	Patch()
	res := Result{}
	Parse([]byte("test"), &res)
	// Output: test
}

func ExampleCheckResult() {
	defer monkey.UnpatchAll()
	Patch()
	CheckResult(failedResponseBytes)
	// Output: RetCode is 1, Reason: failed
}

func ExampleCheckStatusCode() {
	defer monkey.UnpatchAll()
	Patch()
	CheckStatusCode(403)
	// Output: Validation failure.
}

func ExampleCheckStatusCode2() {
	defer monkey.UnpatchAll()
	Patch()
	CheckStatusCode(404)
	// Output: 404 error
}

func ExampleGetResultData() {
	defer monkey.UnpatchAll()
	Patch()
	var data string
	json.Unmarshal(GetResultData(SuccessResponseBytes), &data)
	fmt.Println(data)
	// Output: success
}

func ExampleGetResultStringData() {
	defer monkey.UnpatchAll()
	Patch()
	fmt.Println(GetResultStringData(SuccessResponseBytes))
	// Output: success
}
