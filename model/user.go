package model

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

var err error

type AppUser struct {
	Id        int64     `orm:"column(id)" json:"id"`
	Name      string    `orm:"column(name)" json:"name"`
	Mobile    string    `orm:"column(mobile)" json:"mobile"`
	Status    string    `orm:"column(status)" json:"status"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add"`
}

func (u *AppUser) TableName() string {
	// db table name
	return "app_user"
}

func init() {
	utility.Print(nil, "initializing register AppUser model")
	orm.RegisterModel(new(AppUser))
}

func AddUser(user *AppUser) (id int64, err error) {
	o := orm.NewOrm()
	// default status as "new"
	user.Status = "new"
	id, err = o.Insert(user)
	if err != nil {
		utility.Print(&err, "Err in insert: ", err)
	}
	return
}

func GetAllUsers() ([]AppUser, error) {
	o := orm.NewOrm()
	var users []AppUser
	_, err := o.QueryTable(new(AppUser)).All(&users)
	if err != nil {
		utility.Print(&err, "Err in GetAllUsers: ", err)
		return nil, err
	}
	return users, err
}

func GetUserById(id string) (*AppUser, error) {
	utility.Print(&err, id)
	o := orm.NewOrm()
	var user AppUser
	err := o.QueryTable(new(AppUser)).Filter("id", id).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		utility.Print(&err, "Err in GetUserById: ", err)
		return nil, err
	}
	return &user, err
}

func GetUserByMobile(mob string) (*AppUser, error) {
	utility.Print(&err, mob)
	o := orm.NewOrm()
	var user AppUser
	err := o.QueryTable(new(AppUser)).Filter("mobile", mob).One(&user)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		utility.Print(&err, "Err in GetUserByName: ", err)
		return nil, err
	}
	return &user, err
}

func UpdateUser(user *AppUser) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(user)
	if err != nil {
		utility.Print(&err, "Err in UpdateUser: ", err)
	}
	return
}

func RemoveUser(user *AppUser) (err error) {
	o := orm.NewOrm()
	_, err = o.Delete(user)
	if err != nil {
		utility.Print(&err, "Err in RemoveUser")
	}
	return
}
