package wxmp

import (
	"RemindMe/model"
)

type Dict struct {
	models.Model
    Type  string `json:"type" gorm:"column:type; type:varchar(64); comment:类型名"`
    Id    int    `json:"index" gorm:"column:index; comment:标签编号"`
    Label string `json:"label" gorm:"column:label; type:varchar(64); comment:标签名"`
}

func (Dict) TableName() string {
    return "wxmp_dict"
}

var dict = map[string]map[int]string{
    "重复": {
        0:   "不重复",
        1:   "每天",
        7:   "每周",
        30:  "每月",
        365: "每年",
        355: "农历每年",
    },
    "提醒": {
        -1:    "不提醒",
        0:     "准时",
        5:     "5分钟前",
        10:    "10分钟前",
        15:    "15分钟前",
        30:    "30分钟前",
        60:    "1小时前",
        120:   "2小时前",
        1440:  "1天前",
        2880:  "2天前",
        10080: "1周前",
    },
    "隐私": {
        0: "仅自己可见",
        1: "可分享他人",
        2: "完全公开",
    },
}
