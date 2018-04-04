package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController)  Hello() {
	fmt.Println("33333333")
	//c.StopRun()

	c.Data["json"] = map[string]interface{}{"name": "astaxie"}
	c.ServeJSON()
}
