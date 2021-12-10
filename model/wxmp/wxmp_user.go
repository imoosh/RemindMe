package wxmp

import "RemindMe/global"

type WxmpUser struct {
    global.GVA_MODEL
    Nickname string `json:"nickname"    gorm:"column:nickname; comment:昵称" `
    Gender   int    `json:"gender"      gorm:"column:gender; comment:性别"`
    Avatar   string `json:"avatar"      gorm:"column:avatar; comment:头像"`
    OpenID   string `json:"openid"      gorm:"column:openid; comment:OpenID"`
    UnionID  string `json:"unionid"     gorm:"column:unionid; comment:UnionID"`
    Phone    string `json:"phone"       grom:"column:phone; comment:手机号"`
}
