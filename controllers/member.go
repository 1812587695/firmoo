package controllers

import (
	"github.com/astaxie/beego"
	"firmoo/models"
	"strings"
	"firmoo/utils"
)

type MemberController struct {
	beego.Controller
}

// 查询一条数据
func (c *MemberController)  Get() {

	Id, _ := c.GetInt("id")
	token := c.GetString("token")

	// 验证token
	_,err:=models.MemberTokenToken(token)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"data": "token error"}
	} else {
		res,error := models.MemberOne(Id)
		if error != nil {
			c.Data["json"] = map[string]interface{}{"data": "error"}
		} else {
			c.Data["json"] = map[string]interface{}{"data": res}
		}
	}
	c.ServeJSON()
}

// 登录
func (c *MemberController) Login() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))

	if len(username) == 0 || len(userpwd) == 0 {
		c.Data["json"] = map[string]interface{}{"data": "error "}
	}
	userpwd = utils.String2md5(userpwd)

	members,err := models.MemberLogin(username, userpwd)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"data": "error"}
	} else {
		token_string := Token(username, userpwd)
		token := map[string]interface{}{"token": token_string}

		// 插入member_token表
		err := models.MemberTokenInsert(members.Id, token_string)

		if err != nil {
			c.Data["json"] = map[string]interface{}{"data": "insert error"}
		} else {
			c.Data["json"] = map[string]interface{}{"data": token}
		}
	}

	c.ServeJSON()
}

// 退出登录
func (c *MemberController) LoginOut()  {
	token := c.GetString("token")

	// 验证token
	_,err:=models.MemberTokenToken(token)

	if err != nil {
		c.Data["json"] = map[string]interface{}{"data": "error"}
	}else {
		err = models.MemberTokenDel(token)
		if err == nil {
			c.Data["json"] = map[string]interface{}{"data": "success"}
		}
	}
	c.ServeJSON()
}

// 简单生成token TODO
func Token(username string, userpwd string) string{
	tokens := username + userpwd + "firmoo"
	return utils.String2md5(tokens)
}
