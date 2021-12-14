package system

import (
	"RemindMe/model"
)

type JwtBlacklist struct {
	models.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
