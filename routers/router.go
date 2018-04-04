package routers

import (
	"firmoo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello", &controllers.MainController{}, "Get:Hello")

    beego.Router("/member", &controllers.MemberController{}, "Get:Get")
    beego.Router("/member/login", &controllers.MemberController{}, "Post:Login")
    beego.Router("/member/login_out", &controllers.MemberController{}, "get:LoginOut")

}
