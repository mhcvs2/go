package main

import (
	"os/exec"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}


func main() {
	out, err := exec_shell("ls")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
	slice := strings.Split(out, "\n")
	fmt.Println(slice)
}

/*
bin
lastupdate.tmp
pkg
src

[bin lastupdate.tmp pkg src ]
*/