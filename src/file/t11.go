package main

import (
	"os"
	"fmt"
)

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

func t11() {
	if exist, _ := PathExists("/root/github/go/src/csv/test.csv"); exist {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}

func main() {
	t11()
}