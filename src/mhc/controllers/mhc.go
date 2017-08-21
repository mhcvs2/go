package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"mhc/models"
	"encoding/json"
)

type MhcController struct {
	beego.Controller
}

type Name struct {
	Name string
}

// @Title GetAll
// @Description get all backup
// @Success 200 {object} models.BackupList
// @router / [get]
func (this *MhcController) GetAll() {

	this.Data["json"] = models.BackupGetAll()
	fmt.Println(this.Data["json"])

	this.ServeJSON()
}

// @Title Trigger Backup
// @Description trigger backup
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *MhcController) Post() {
	var name Name
	fmt.Println(u.Ctx.Input.RequestBody)
	json.Unmarshal(u.Ctx.Input.RequestBody, &name)
	fmt.Println(name)
	//name.Name = u.GetString("name")
	//fmt.Println(u.Input())
	u.Data["json"] = models.TriggerBackup(name.Name)
	u.ServeJSON()
}
