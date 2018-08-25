package proxy

import (
	"designPattern/AAF_proxy/a_common/gamePlayer"
	"fmt"
)

type playerProxy struct {
	player gamePlayer.IPlayer
}

func NewPlayerProxy(name string) *playerProxy {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err: ", err)
		}
	}()
	pp := playerProxy{}
	pp.player = gamePlayer.NewPlayer(&pp, name)
	return &pp
}

func (p *playerProxy)Login(user, password string) {
	p.player.Login(user, password)
}

func (p *playerProxy)KillBoss() {
	p.player.KillBoss()
}

func (p *playerProxy)Upgrade() {
	p.player.Upgrade()
}


