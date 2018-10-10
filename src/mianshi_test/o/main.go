package main

import (
	"fmt"
	"errors"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func DoTheThing2(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err2 := tryTheThing()
		if err2 != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func DoTheThing3(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string,error)  {
	return "",ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
	fmt.Println(DoTheThing2(true))
	fmt.Println(DoTheThing3(true))
}
//<nil>
//<nil>
//did not work
//did not work