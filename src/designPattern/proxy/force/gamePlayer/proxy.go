package gamePlayer

type playerProxy struct {
	player IPlayer
}

func NewPlayerProxy(player IPlayer) *playerProxy {
	pp := playerProxy{player:player}
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

func (p *playerProxy)GetProxy() IPlayer {
	return p
}


