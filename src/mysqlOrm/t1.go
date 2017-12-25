package main

//引用模块
import "github.com/ablegao/orm"

//mysql 驱动
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func init() {
	//建立连接
	// 参数分别为 名称 ， 驱动， 连接字符串
	// 注：必须包含一个default 连接， 作为默认连接。
	orm.NewDatabase("default", "mysql", "root:root.123@tcp(109.105.4.65:32981)/ormtest?charset=utf8")
}

//CREATE TABLE IF NOT EXISTS `user_info` (  `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',   `username` VARCHAR(45) NOT NULL COMMENT '用户名称', `password` VARCHAR(45) NOT NULL COMMENT '用户密码',   PRIMARY KEY (`id`))  ENGINE = InnoDB  AUTO_INCREMENT = 1  DEFAULT CHARACTER SET = utf8  COLLATE = utf8_general_ci  COMMENT = '用户表';

//建立一个数据模型。
type UserInfo struct {
	orm.Object
	Id     int64  `field:"id" auto:"true" index:"pk"`
	Name   string `field:"username"`
	Passwd string `field:"password"`
}

//数据库表名称
func (self *UserInfo) GetTableName() string {
	return "ormtest.user_info"
}

func Search(UserName string) *UserInfo {
	//查询一个用户名为 "test1"的账户
	user := new(UserInfo)
	user.Objects(user).Filter("Name", UserName).One()
	fmt.Println(user.Id, user.Passwd, user.Name)
	return user
}

func Update(UserName, newName string) {
	user := Search(UserName)
	//Update
	user.Name = "test2"
	user.Objects(user).Save()
	// or
	user.Objects(user).Filter("Id", 1).Change("Name", "test2").Save()
}

func SearchID() {
	//查询id小于10的所有数据
	user := new(UserInfo)
	users, err := user.Objects(user).Filter("Id__lt", 10).All()
	if err == nil {
		for _, userinfo := range users {
			u := userinfo.(*UserInfo)
			fmt.Println(u.Id, u.Passwd, u.Name)
		}
	}
}

func Create(name string) error {
	//Create
	user := new(UserInfo)
	user.Name = name
	user.Passwd = "123456"
	_, _, err := user.Objects(user).Save()
	return err
}

////delete
//user.Objects(user).Delete()
//
//// User other Database connect
//orm.NewDatabase("other" , "mysql" , "user:passwd@ip/database?charset=utf8")
//user.Objects(user).Db("other").Filter(x ,x).Delete()
//// or
//user.Objects(user).Filter().Db("other").XXX()

func main() {
	for i := 1; i<200; i++ {
		Create("test" + strconv.Itoa(i))
	}
}
