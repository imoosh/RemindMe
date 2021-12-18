package wxmp

import (
    "RemindMe/model"
)

type WxmpActivitySubscribers struct {
    models.Model
    ActivityId       uint `json:"activity_id"        gorm:"primaryKey;autoIncrement:false; column:activity_id;index:index_activity_id;comment:活动id;"`
    SubscriberUserId uint `json:"subscriber_user_id" gorm:"primaryKey;autoIncrement:false; column:subscriber_user_id; comment:订阅者用户id"`
    RemindAt         int  `json:"remind_at"          gorm:"column:remind_at; comment:订阅者个人的提醒时间,见字典项"`
    IsTplRemind      bool `json:"is_tpl_remind"      gorm:"column:is_tpl_remind; default:0; comment:模板提醒"`
    IsSmsRemind      bool `json:"is_sms_remind"      gorm:"column:is_sms_remind; default:0; comment:短信提醒"`
    Status           int  `json:"status"             gorm:"column:status; type:int(2);default:0; comment:订阅状态,0:订阅取消，1:订阅成功"`
}
