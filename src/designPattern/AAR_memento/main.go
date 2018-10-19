package main

import "fmt"

type IState interface {
	GetState() string
	SetState(state string)
}

////////////////////////////////////////

type Memento struct {
	state string
}

func NewMemento(state string) *Memento {
	return &Memento{state:state}
}

func (m *Memento) GetState() string {
	return m.state
}

func (m *Memento) SetState(state string) {
	m.state = state
}

////////////////////////////////////////////

type Boy struct {
	state string
}

func NewBoy() *Boy {
	return &Boy{state:""}
}

func (m *Boy) GetState() string {
	return m.state
}

func (m *Boy) SetState(state string) {
	m.state = state
}

func (m *Boy) CreateMemento() IState {
	return NewMemento(m.state)
}

func (m *Boy) RestoreMemento(memento IState) {
	m.state = memento.GetState()
}
////////////////////////////////////////////////

type Caretaker struct {
	memento IState
}

func NewCaretaker() *Caretaker {
	return new(Caretaker)
}

func (c *Caretaker) GetMemento() IState {
	return c.memento
}

func (c *Caretaker) SetMemento(memento IState) {
	c.memento = memento
}

/////////////////////////////////////////////////////

func main() {
	b := NewBoy()
	c := NewCaretaker()
	b.SetState("good")
	fmt.Println(b.GetState())
	c.SetMemento(b.CreateMemento())
	b.SetState("bad")
	fmt.Println(b.GetState())
	b.RestoreMemento(c.GetMemento())
	fmt.Println(b.GetState())
}

//good
//bad
//good