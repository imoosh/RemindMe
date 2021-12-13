package wxmp

import "RemindMe/global"

type WxmpActivitySubscribers struct {
    global.GVA_MODEL
    ActivityId       int64  `json:"activity_id" gorm:"column:activity_id; comment:活动id; index: index_activity_id"`
    SubscriberUserId string `json:"subscriber_user_id" gorm:"column:subscriber_user_id; comment:订阅者用户id"`

    RemindAt    int  `json:"remind_at" gorm:"column:remind_at; comment:订阅者个人的提醒时间,见字典项"`
    IsTplRemind bool `json:"is_tpl_remind" gorm:"column:is_tpl_remind; comment:模板提醒"`
    IsSmsRemind bool `json:"is_sms_remind" gorm:"column:is_sms_remind; comment:短信提醒"`
}
