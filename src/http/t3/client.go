package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
)

func httpGet(p string) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3322/mhc/%s", p))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(body))
}


func main() {
	httpGet("2")
}