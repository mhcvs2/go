package main

import "designPattern/AAF_proxy/a_common/proxy"

func main() {
	p := proxy.NewPlayerProxy("mhc")
	p.Login("mhcuser", "mhc.123")
	p.KillBoss()
	p.Upgrade()
}
