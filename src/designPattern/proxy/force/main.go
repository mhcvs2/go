package main

import "designPattern/proxy/force/gamePlayer"

func main() {
	p := gamePlayer.NewPlayer("mhc")
	p.Login("mhcuser", "mhc.123")
	p.KillBoss()
	p.Upgrade()

	q := p.GetProxy()
	q.Login("mhcuser", "mhc.123")
	q.KillBoss()
	q.Upgrade()
}
