package main

import (
	"net"
	"fmt"
	"os"
)

func t11() {
	name := "192.168.1.97"      //只能是ip, 不能加端口
	ip := net.ParseIP(name)

	if ip == nil {
		fmt.Fprint(os.Stderr, "err: 无效的地址")
	}

	fmt.Printf("ip: %s\n", ip.String())

	defaultMask := ip.DefaultMask()
	fmt.Printf("defaultMask: %s, %s\n", defaultMask, defaultMask.String())

	ones, bits := defaultMask.Size()
	fmt.Printf("ones: %d, bits: %d\n", ones, bits)
}
//ip: 192.168.1.97
//defaultMask: ffffff00, ffffff00
//ones: 24, bits: 32

func t12() {
	name := "192.168.1.97"

	ip := net.ParseIP(name)

	mask := ip.DefaultMask()

	network := ip.Mask(mask)

	fmt.Fprintf(os.Stdout, "network: %s", network.String()) // 192.168.1.0
}
//network: 192.168.1.0

func t13() {
	domain := "www.baidu.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "ipAddr: %s, IP: %s, Network: %s, Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "%s\n", n) // 115.239.210.26    115.239.210.27 这2个地址打开都是百度
	}
}
//ipAddr: 220.181.112.244, IP: 220.181.112.244, Network: ip, Zone:
//220.181.112.244
//220.181.111.188

func main() {
	t13()
}
