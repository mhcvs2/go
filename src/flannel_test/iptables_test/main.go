package main

import (
	"github.com/coreos/flannel/network"
	"github.com/coreos/flannel/pkg/ip"
	"github.com/coreos/flannel/subnet"
)

func t1() {
	subnetIP, _ := ip.ParseIP4("1.2.3.1")
	n := ip.IP4Net{IP: subnetIP, PrefixLen: 16}
	sn := ip.IP4Net{IP: subnetIP, PrefixLen: 24}
	l := &subnet.Lease{
		Subnet: sn,
	}
	rules := network.MasqRules(n, l)
	network.SetupAndEnsureIPTables(rules, 10)
}

func t2() {
	subnetIP, _ := ip.ParseIP4("1.2.3.1")
	n := ip.IP4Net{IP: subnetIP, PrefixLen: 16}
	sn := ip.IP4Net{IP: subnetIP, PrefixLen: 24}
	l := &subnet.Lease{
		Subnet: sn,
	}
	rules := network.MasqRules(n, l)
	network.DeleteIPTables(rules)
}

func t3() {
	network.SetupAndEnsureIPTables(network.ForwardRules("1.2.0.0/16"), 10)
}

func t4() {
	network.DeleteIPTables(network.ForwardRules("1.2.0.0/16"))
}

func main() {
	t3()
}
