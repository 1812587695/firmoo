package models

import (
	"github.com/astaxie/beego/orm"
)

func (a *MemberToken) TableName() string {
	return TableName("member_token")
}

type MemberToken struct {
	Id       int
	MemberId int
	Token    string
}


func MemberTokenOne(member_id int) (*MemberToken, error) {
	o := orm.NewOrm()
	m := MemberToken{MemberId: member_id}
	err := o.Read(&m, "MemberId")  // 通过member_id查询
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func MemberTokenToken(token string) (*MemberToken, error) {
	o := orm.NewOrm()
	m := MemberToken{Token: token}
	err := o.Read(&m, "Token") // 通过token查询
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// insert
func MemberTokenInsert(member_id int, token string) (error) {

	_, err := MemberTokenOne(member_id)

	o := orm.NewOrm()
	if err != nil{
		member_token := MemberToken{MemberId: member_id, Token:token}
		// 增加
		_, err = o.Insert(&member_token)
	} else {

		_, err = o.Raw("UPDATE "+ MemberTokenTBName() +" SET token = ? where member_id = ?", token, member_id).Exec()
	}

	return err
}

func MemberTokenDel(token string) error {

	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM "+ MemberTokenTBName() +" WHERE (`token`=?)", token).Exec()

	return err
}


func init() {
	// 注册model
	orm.RegisterModel(new(MemberToken))
}
