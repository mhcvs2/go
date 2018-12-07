package main

import (
	"github.com/asaskevich/govalidator"
	"fmt"
)

func init() {
	govalidator.TagMap["string_not_nil"] = govalidator.Validator(
		func(str string) bool {
			return str != "" && str != "0"
		},
	)
}

type SignInfo struct {
	ID          string    `valid:"string_not_nil,required"`
	Location    string
	Subject     string
	PostAddress string
}

type SignInfo4Pool struct {
	SignInfo
	Key string
}

var pool = make(map[string]*SignInfo4Pool)

func GetSignInfo(key string) *SignInfo4Pool {
	if si, ok := pool[key]; ok {
		fmt.Println("get in pool")
		return si
	} else {
		res := new(SignInfo4Pool)
		pool[key] = res
		fmt.Printf("put inf pool: %s\n", key)
		return res
	}
}

func main() {
	for i:=0; i<4; i++ {
		subject := fmt.Sprintf("subject-%d", i)
		for j:=0; j<30; j++ {
			key := fmt.Sprintf("%s-location-%d", subject, j)
			GetSignInfo(key)
		}
	}
	signInfo := GetSignInfo("subject-1-location-1")
	signInfo.ID = "0"
	if _, err := govalidator.ValidateStruct(signInfo); err != nil {
		panic(err)
	}
	fmt.Println(signInfo)
}

//put inf pool: subject-3-location-28
//put inf pool: subject-3-location-29
//get in pool
//panic: ID: 0 does not validate as string_not_nil
//
//goroutine 1 [running]:
//main.main()
///root/github/go/src/designPattern/flyweight/main.go:53 +0x30e
//exit status 2
