package main

import "fmt"

type IProject interface {
	Add(name string, num, cost int)
	GetProjectInfo()
	Iterator() IProjectIterator
}

type IProjectIterator interface {
	HasNext() bool
	Next() IProject
}

///////////////////////////////////////////

type Project struct {
	name string
	num int
	cost int
	projectList []IProject
}

func NewProject(name string, num, cost int) *Project {
	p := &Project{name:name, num:num, cost:cost}
	p.projectList = append(p.projectList, p)
	return p
}

func (p *Project) Add(name string, num, cost int) {
	p.projectList = append(p.projectList, NewProject(name, num, cost))
}

func (p *Project) GetProjectInfo(){
	fmt.Println("name is", p.name)
	fmt.Println("num is", p.num)
	fmt.Println("cost is", p.cost)
}

func (p *Project) Iterator() IProjectIterator {
	return NewProjectIterator(p.projectList)
}

/////////////////////////////////////////////////////////

type ProjectIterator struct {
	projectList []IProject
	current int
}

func NewProjectIterator(projectList []IProject) *ProjectIterator {
	return &ProjectIterator{projectList, 0}
}

func (p *ProjectIterator) HasNext() bool {
	return p.current < len(p.projectList) && p.projectList[p.current] != nil
}

func (p *ProjectIterator) Next() IProject {
	res := p.projectList[p.current]
	p.current++
	return res
}

func main() {
	p := NewProject("p1", 1, 111)
	p.Add("p4", 4, 444)
	p.Add("p2", 2, 222)
	p.Add("p3", 3, 333)
	for it := p.Iterator(); it.HasNext();{
		it.Next().GetProjectInfo()
	}
}

//name is p1
//num is 1
//cost is 111
//name is p4
//num is 4
//cost is 444
//name is p2
//num is 2
//cost is 222
//name is p3
//num is 3
//cost is 333