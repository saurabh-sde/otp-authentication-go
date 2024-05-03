package model

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

type LogEvent struct {
	Id        int64     `orm:"column(id)" json:"id"`
	EventName string    `orm:"column(event_name)" json:"name"`
	UserId    string    `orm:"column(log_id)" json:"mobile"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add"`
}

func (u *LogEvent) TableName() string {
	// db table name
	return "log_event"
}

func init() {
	utility.Print(nil, "initializing register LogEvent model")
	orm.RegisterModel(new(LogEvent))
}

func AddLogEvent(log *LogEvent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(log)
	if err != nil {
		utility.Print(&err, "Err in insert: ", err)
	}
	return
}

func GetAllLogEvents() ([]LogEvent, error) {
	o := orm.NewOrm()
	var logs []LogEvent
	_, err := o.QueryTable(new(LogEvent)).All(&logs)
	if err != nil {
		utility.Print(&err, "Err in GetAllLogEvents: ", err)
		return nil, err
	}
	return logs, err
}

func GetLogEventByUserId(id string) (*LogEvent, error) {
	utility.Print(&err, id)
	o := orm.NewOrm()
	var log LogEvent
	err := o.QueryTable(new(LogEvent)).Filter("user_id", id).One(&log)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, nil
		}
		utility.Print(&err, "Err in GetLogEventById: ", err)
		return nil, err
	}
	return &log, err
}
