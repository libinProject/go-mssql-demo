package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (self *MainController) GetTable() {
	out := make(map[string]interface{})
	out["code"] = 123
	out["msg"] = 123
	out["count"] = 1
	out["data"] = models.GetPageList()
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
