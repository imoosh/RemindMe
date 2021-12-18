package wxmp

import (
    "RemindMe/model"
)

type WxmpUser struct {
    models.Model
    Nickname   string          `json:"nickname"    gorm:"column:nickname; type:varchar(64); comment:昵称" `
    Gender     int             `json:"gender"      gorm:"column:gender; comment:性别"`
    Avatar     string          `json:"avatar"      gorm:"column:avatar; type:varchar(256); comment:头像"`
    OpenID     string          `json:"openid"      gorm:"column:openid; type:varchar(128); comment:OpenID; uniqueIndex:idx_openid"`
    UnionID    string          `json:"unionid"     gorm:"column:unionid; type:varchar(128); comment:UnionID"`
    Phone      string          `json:"phone"       grom:"column:phone; type:varchar(16); comment:手机号"`
    Activities []*WxmpActivity `gorm:"many2many:user_activity;"`
}
