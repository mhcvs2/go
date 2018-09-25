package corp

import "fmt"

type ICorp interface {
	GetInfo() string
	SetParent(corp ICorp)
	GetParent() ICorp
}

type IBranch interface {
	AddSubordinate(corp ICorp)
	GetSubordinateInfo() []ICorp
}

//---------------------------------------------------------------

type Corp struct {
	name string
	position string
	salary int
	Parent ICorp
}

func NewCorp(name, position string, salary int) *Corp {
	return &Corp{
		name: name,
		position: position,
		salary: salary,
	}
}

func (c *Corp) GetInfo() string {
	return fmt.Sprintf("name: %s, position: %s, salary: %d", c.name, c.position, c.salary)
}

func (c *Corp) SetParent(corp ICorp) {
	c.Parent = corp
}

func (c *Corp) GetParent() ICorp {
	return c.Parent
}

//------------------------------------------------------------------------------

type Leaf struct {
	Corp
}

func NewLeaf(name, position string, salary int) *Leaf {
	return &Leaf{
		 *NewCorp(name, position, salary),
	}
}

//------------------------------------------------------------------------------

type Branch struct {
	Corp
	SubordinateList []ICorp
}

func NewBranch(name, position string, salary int) *Branch {
	return &Branch{
		*NewCorp(name, position, salary),
		[]ICorp{},
	}
}

func (b *Branch) AddSubordinate(corp ICorp) {
	corp.SetParent(b)
	b.SubordinateList = append(b.SubordinateList, corp)
}

func (b *Branch) GetSubordinateInfo() []ICorp {
	return b.SubordinateList
}

