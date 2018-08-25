package gamePlayer

import "fmt"

type IPlayer interface {
	Login(user, password string)
	KillBoss()
	Upgrade()
}

type player struct {
	Name string
}

func NewPlayer(proxy IPlayer, name string) *player {
	if proxy == nil {
		panic("can not create real player")
	}
	return &player{Name:name}
}

func (p *player)Login(user, password string) {
	fmt.Println("user: ", user, "password:", password)
}

func (p *player)KillBoss() {
	fmt.Println(p.Name, "kill boss")
}

func (p *player)Upgrade() {
	fmt.Println(p.Name, "Upgrade")
}

