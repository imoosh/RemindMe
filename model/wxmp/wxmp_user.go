package wxmp

import "RemindMe/global"

type WxmpUser struct {
    global.GVA_MODEL
    Nickname string `json:"nickname"    gorm:"column:nickname; type:varchar(64); comment:昵称" `
    Gender   int    `json:"gender"      gorm:"column:gender; comment:性别"`
    Avatar   string `json:"avatar"      gorm:"column:avatar; type:varchar(128); comment:头像"`
    OpenID   string `json:"openid"      gorm:"column:openid; type:varchar(128); comment:OpenID"`
    UnionID  string `json:"unionid"     gorm:"column:unionid; type:varchar(128); comment:UnionID"`
    Phone    string `json:"phone"       grom:"column:phone; type:varchar(16); comment:手机号"`
}
