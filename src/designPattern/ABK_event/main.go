package main

import (
	"designPattern/ABK_event/product"
	"fmt"
)

func main() {
	dispatch := product.GetEventDispatch()
	dispatch.RegisterCustomer(product.NewBeggar())
	dispatch.RegisterCustomer(product.NewCommoner())
	dispatch.RegisterCustomer(product.NewNobleman())

	manager := product.NewProductManager()
	p := manager.CreateProduct("小男孩")
	manager.EditProduct(p, "hehe")
	p2 := manager.Clone(p)
	fmt.Println(p2.Name())
	manager.AbandonProduct(p)

}
//小男孩
//平民处理事件: 小男孩 新建, 事件类型: 0
//小男孩
//贵族处理事件: 小男孩 修改, 事件类型: 2
//hehe
//hehe
//乞丐处理事件: hehe 销毁, 事件类型: 1