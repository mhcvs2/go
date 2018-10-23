package employee1

import (
	"fmt"
	"reflect"
)

type IVisitor interface {
	Visit(employee interface{})
}

type Visitor struct {
}

func (v *Visitor)getBasicInfo(employee IEmployee) string {
	return fmt.Sprintf("name is : %s", employee.GetName())
}


func (v *Visitor)getManagerInfo(manager IManager) string {
	return fmt.Sprintf("%s\tper: %s", v.getBasicInfo(manager), manager.GetPer())
}

func (v *Visitor)getCommonEmployeeInfo(ce ICommonEmployee) string {
	return fmt.Sprintf("%s\tjob: %s", v.getBasicInfo(ce), ce.GetJob())
}

func (v *Visitor)Visit(employee interface{}){
	if reflect.TypeOf(employee) == reflect.TypeOf(new(commonEmployee)){
		fmt.Println(v.getCommonEmployeeInfo(employee.(*commonEmployee)))
	} else if reflect.TypeOf(employee) == reflect.TypeOf(new(manager)){
		fmt.Println(v.getManagerInfo(employee.(*manager)))
	} else {
		fmt.Println("unknown type")
	}
}

