package main

import (
	"fmt"
	"github.com/coreos/flannel/pkg/ip"
)

func t1() {
	iface, err := ip.GetDefaultGatewayIface()
	if err == nil {
		fmt.Println(iface.Name)
	} else {
		panic(err)
	}
}

func main() {
	t1()
}
