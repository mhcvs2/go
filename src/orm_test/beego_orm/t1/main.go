package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `orm:"auto"`
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(ali:3306)/gorm_test?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "mhc"}

	if id, err := o.Insert(&user); err != nil {
		fmt.Println("insert failed, err: ", err.Error())
	} else {
		fmt.Println("insert success, id: ", id)
	}
	user.Name = "astaxie"
	num, _ := o.Update(&user)
	fmt.Println(num)

	u := User{Id: user.Id}
	o.Read(&u)
	fmt.Println(u.Name)

	num, _ = o.Delete(&u)
}
