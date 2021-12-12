package wxmp

import (
	"RemindMe/global"
	"time"
)

type WxmpActivity struct {
	global.GVA_MODEL
	Title     string `json:"title" gorm:"column:title; type:varchar(64); comment:活动名称"`
	Address   string `json:"address" gorm:"column:address; type:varchar(128); comment:活动地址"`
	Latitude  string `json:"latitude" gorm:"column:latitude; type:varchar(128); comment:地址纬度"`
	Longitude string `json:"longitude" gorm:"column:longitude; type:varchar(128); comment:地址精度"`
	NotificationTime time.Time `json:"notificationTime" gorm:"column:"`

	Nickname string `json:"nickname"    gorm:"column:nickname; comment:昵称" `
	Gender   int    `json:"gender"      gorm:"column:gender; comment:性别"`
	Avatar   string `json:"avatar"      gorm:"column:avatar; comment:头像"`
	OpenID   string `json:"openid"      gorm:"column:openid; comment:OpenID"`
	UnionID  string `json:"unionid"     gorm:"column:unionid; comment:UnionID"`
	Phone    string `json:"phone"       grom:"column:phone; comment:手机号"`
}
