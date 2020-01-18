package model

import "time"

const (
	SexWomen = "W"
	SexMan = "M"
	SexUnknown = "U"
)

type User struct {
	Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string 		`xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd       string	`xorm:"varchar(40)" form:"passwd" json:"-"`   // 用户密码 md5(passwd + salt)
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex        string	`xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt       string	`xorm:"varchar(10)" form:"salt" json:"-"`
	Online     int	`xorm:"int(10)" form:"online" json:"online"`   //是否在线
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   //用户鉴权
	Memo      string	`xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat   time.Time	`xorm:"datetime" form:"createat" json:"createat"`   //创建时间, 统计用户增量时使用
}

