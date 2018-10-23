package visitor

type Role interface {
	Accept(actor IActor)
}

type IdiotRole struct {
}

func (r IdiotRole) Accept(actor IActor) {
	actor.Act(r)
}

type KungFuRole struct {
}

func (r KungFuRole) Accept(actor IActor) {
	actor.Act(r)
}