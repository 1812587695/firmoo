package main

import (
	_ "firmoo/routers"
	_ "firmoo/sysinit"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	beego.Run()
}

