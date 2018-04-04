package models

import (
	"github.com/astaxie/beego"
)

////初始化
//func init() {
//	orm.RegisterModel(new(Member))
//}

//下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

//获取对应的表名称
func MemberTBName() string {
	return TableName("member")
}

func MemberTokenTBName() string {
	return TableName("member_token")
}

