package employee

type IEmployee interface {
	GetName() string
	SetName(name string)
	Accept(visitor IVisitor)
}

////////////////////////////////////////
type employee struct {
	name string
}

func NewEmployee() *employee {
	return new(employee)
}

func (e *employee)GetName() string {
	return e.name
}

func (e *employee)SetName(name string) {
	e.name = name
}

func (e *employee)Accept(visitor IVisitor) {
	panic("not implement")
}

/////////////////////////////////////////////

type ICommonEmployee interface {
	IEmployee
	GetJob() string
	SetJob(job string)
}

type commonEmployee struct {
	employee
	job string
}

func NewCommonEmployee() *commonEmployee {
	return new(commonEmployee)
}

func (c *commonEmployee) GetJob() string {
	return c.job
}

func (c *commonEmployee) SetJob(job string) {
	c.job = job
}

func (c *commonEmployee) Accept(visitor IVisitor) {
	visitor.Visit(c)
}

///////////////////////////////////////////////////////

type IManager interface {
	IEmployee
	GetPer() string
	SetPer(job string)
}

type manager struct {
	employee
	per string
}

func NewManager() *manager {
	return new(manager)
}

func (c *manager) GetPer() string {
	return c.per
}

func (c *manager) SetPer(per string) {
	c.per = per
}

func (c *manager) Accept(visitor IVisitor) {
	visitor.Visit(c)
}
