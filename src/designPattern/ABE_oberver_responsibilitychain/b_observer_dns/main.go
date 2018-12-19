package main

import (
	"designPattern/ABE_oberver_responsibilitychain/b_observer_dns/dns"
	"designPattern/ABE_oberver_responsibilitychain/b_observer_dns/observer"
	"fmt"
)

func main() {
	sh := dns.NewSHDnsServer()
	china := dns.NewChinaDnsServer()
	top := dns.NewGlobalDnsServer()
	china.SetUpperServer(top)
	sh.SetUpperServer(china)

	recoder := new(dns.Recorder)
	recoder.Domain = "mm.sh.cna"
	sh.Update(observer.NewObservable(), recoder)
	fmt.Println(recoder.Ip)
	fmt.Println(recoder.Owner)
}
