package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id int `orm:"auto"`
	Name string `orm:"size(100)"`
}


func t1() {
	orm.RegisterModel(new(User))
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(ali:3306)/gorm_test?charset=utf8", 30)

	// create table
	orm.RunSyncdb("default", false, true)
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

type User2 struct {
	Id int `orm:"auto"`
	Name string `orm:"size(100)"`
	Age  int
}

func (u *User2) TableName() string {
	return "user"
}

func t2() {
	orm.RegisterModel(new(User2))
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(ali:3306)/db_test?charset=utf8", 30)
	o := orm.NewOrm()

	u := User2{Name: "haha", Age: 99}
	if num, err := o.Insert(&u); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
}

func t3() {
	orm.RegisterModel(new(User2))
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(ali:3306)/db_test?charset=utf8", 30)
	o := orm.NewOrm()

	u := User2{Id: 5}
	o.Read(&u)
	fmt.Println("u.Name: ", u.Name)

	u2 := User2{}
	o.QueryTable("user").Filter("age", 88).One(&u2)
	fmt.Println("u.Name: ", u2.Name)
}

func main() {
	t3()
}
