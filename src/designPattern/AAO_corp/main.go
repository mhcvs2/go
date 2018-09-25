package main

import (
	"designPattern/AAO_corp/corp"
	"bytes"
	"reflect"
	"fmt"
)

func CompositeCorpTree() *corp.Branch {
	ceo := corp.NewBranch("王大麻子", "总经理", 100000)

	developDep := corp.NewBranch("刘大瘸子","研发部门经理",20000)
	salesDep := corp.NewBranch("马二拐子","销售部门经理",20000)
	financeDep := corp.NewBranch("赵三坨子","财务部经理",30000)

	firstDevGroup := corp.NewBranch("杨三奥迪","开发一组组长",5000)
	secondDevGroup := corp.NewBranch("吴大棒槌","开发二组组长",6000)

	a := corp.NewLeaf("a", "研发人员", 2000)
	b := corp.NewLeaf("b", "研发人员", 2000)
	c := corp.NewLeaf("c", "研发人员", 2000)
	d := corp.NewLeaf("d", "研发人员", 2000)
	e := corp.NewLeaf("e", "研发人员", 2000)
	f := corp.NewLeaf("f", "研发人员", 2000)
	h := corp.NewLeaf("h", "销售人员", 2000)
	i := corp.NewLeaf("i", "销售人员", 2000)
	j := corp.NewLeaf("j", "财务人员", 2000)
	k := corp.NewLeaf("k", "ceo秘书", 2000)
	zhengLaoLiu := corp.NewLeaf("郑老六", "研发部副总", 2000)

	ceo.AddSubordinate(developDep)
	ceo.AddSubordinate(salesDep)
	ceo.AddSubordinate(financeDep)
	ceo.AddSubordinate(k)

	developDep.AddSubordinate(firstDevGroup)
	developDep.AddSubordinate(secondDevGroup)
	developDep.AddSubordinate(zhengLaoLiu)

	firstDevGroup.AddSubordinate(a)
	firstDevGroup.AddSubordinate(b)
	firstDevGroup.AddSubordinate(c)
	secondDevGroup.AddSubordinate(d)
	secondDevGroup.AddSubordinate(e)
	secondDevGroup.AddSubordinate(f)

	salesDep.AddSubordinate(h)
	salesDep.AddSubordinate(i)
	financeDep.AddSubordinate(j)
	return ceo
}

func GetStructName(s interface{}) string {
	rootValue := reflect.ValueOf(s)
	if rootValue.Kind() == reflect.Ptr {
		rootValue = rootValue.Elem()
	}
	return rootValue.Type().Name()
}

func GetTreeInfo(root interface{}) string {
	subordinateList := root.(*corp.Branch).GetSubordinateInfo()
	info := new(bytes.Buffer)
	for _, s := range subordinateList {
		if GetStructName(s) == "Leaf" {
			info.WriteString(s.GetInfo())
			info.WriteString("\n")
		} else {
			info.WriteString(s.GetInfo())
			info.WriteString("\n")
			info.WriteString(GetTreeInfo(s))
		}
	}
	return info.String()
}

func main() {
	ceo := CompositeCorpTree()
	fmt.Println(ceo.GetInfo())
	fmt.Println(GetTreeInfo(ceo))
}

//name: 王大麻子, position: 总经理, salary: 100000
//name: 刘大瘸子, position: 研发部门经理, salary: 20000
//name: 杨三奥迪, position: 开发一组组长, salary: 5000
//name: a, position: 研发人员, salary: 2000
//name: b, position: 研发人员, salary: 2000
//name: c, position: 研发人员, salary: 2000
//name: 吴大棒槌, position: 开发二组组长, salary: 6000
//name: d, position: 研发人员, salary: 2000
//name: e, position: 研发人员, salary: 2000
//name: f, position: 研发人员, salary: 2000
//name: 郑老六, position: 研发部副总, salary: 2000
//name: 马二拐子, position: 销售部门经理, salary: 20000
//name: h, position: 销售人员, salary: 2000
//name: i, position: 销售人员, salary: 2000
//name: 赵三坨子, position: 财务部经理, salary: 30000
//name: j, position: 财务人员, salary: 2000
//name: k, position: ceo秘书, salary: 2000