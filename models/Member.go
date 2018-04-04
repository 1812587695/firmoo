package models

import (
	"github.com/astaxie/beego/orm"
)

func (a *Member) TableName() string {
	return MemberTBName()
}

type Member struct {
	Id                 int
	RealName           string `orm:"size(32)"`
	UserName           string `orm:"size(24)"`
	UserPwd            string `json:"-"`
	IsSuper            bool
	Status             int
	Mobile             string                `orm:"size(16)"`
	Email              string                `orm:"size(256)"`
	Avatar             string                `orm:"size(256)"`
}

// 根据id获取单条
func MemberOne(id int) (*Member, error) {
	o := orm.NewOrm()
	m := Member{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func MemberLogin(username string, userpwd string) (*Member, error) {


	m := Member{}
	err := orm.NewOrm().QueryTable(MemberTBName()).Filter("username", username).Filter("userpwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func init() {
	// 注册model
	orm.RegisterModel(new(Member))
}