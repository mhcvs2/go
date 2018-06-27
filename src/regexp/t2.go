package main

import (
	"regexp"
	"fmt"
)

var text = `sdfsdfinfra.gcloud.srcb.com`

var target = `gcloud.srcb.com`

func t21() {
	reg := regexp.MustCompile(fmt.Sprintf("%s$", target))
	if reg.MatchString(text) {
		reg = regexp.MustCompile(fmt.Sprintf(".%s", target))
		fmt.Println(reg.ReplaceAllString(text, ""))
	} else {
		fmt.Println("ignore")
	}
	//true
}

func main() {
	t21()
}
