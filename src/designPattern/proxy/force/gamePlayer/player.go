package gamePlayer

import (
	"fmt"
)

type IPlayer interface {
	Login(user, password string)
	KillBoss()
	Upgrade()
	GetProxy() IPlayer
}

type player struct {
	proxy IPlayer
	Name string
}

func NewPlayer(name string) *player {
	return &player{Name:name}
}

func (p *player)Login(user, password string) {
	if !p.isProxy() {
		fmt.Println("please set proxy and use proxy to call")
		return
	}
	fmt.Println("user: ", user, "password:", password)
}

func (p *player)KillBoss() {
	if !p.isProxy() {
		fmt.Println("please set proxy and use proxy to call")
		return
	}
	fmt.Println(p.Name, "kill boss")
}

func (p *player)Upgrade() {
	if !p.isProxy() {
		fmt.Println("please set proxy and use proxy to call")
		return
	}
	fmt.Println(p.Name, "Upgrade")
}

func (p *player)GetProxy() IPlayer {
	p.proxy = NewPlayerProxy(p)
	return p.proxy
}

func (p *player)isProxy() bool {
	return p.proxy != nil
}
