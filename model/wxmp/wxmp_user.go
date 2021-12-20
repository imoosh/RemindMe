package wxmp

import (
	"RemindMe/model"
)

type User struct {
	models.Model
	Nickname      string `gorm:"column:nickname; type:varchar(64); comment:昵称" `
	Gender        int    `gorm:"column:gender; comment:性别"`
	Avatar        string `gorm:"column:avatar; type:varchar(256); comment:头像"`
	OpenID        string `gorm:"column:openid; type:varchar(128); comment:OpenID; uniqueIndex:idx_openid"`
	UnionID       string `gorm:"column:unionid; type:varchar(128); comment:UnionID"`
	Phone         string `gorm:"column:phone; type:varchar(16); comment:手机号"`
	Activities    []Activity `gorm:"foreignKey:PublisherID"`
	Subscriptions []Activity `gorm:"many2many:wxmp_activity_subscription"`
}

func (User) TableName() string {
	return "wxmp_user"
}
