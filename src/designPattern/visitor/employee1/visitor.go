package employee1

import "fmt"

type IVisitor interface {
	Visit(employee IEmployee)
}

type Visitor struct {

}

func (v *Visitor)Visit(employee IEmployee) {
	fmt.Println(employee.GetName())
}

