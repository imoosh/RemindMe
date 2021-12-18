package wxmp

import (
	"RemindMe/global"
	"RemindMe/model"
	"go.uber.org/zap"
)

type ActivitySubscriber struct {
	models.Model
	ActivityId uint `gorm:"primaryKey;autoIncrement:false; index:index_activity_id;comment:活动id;"`
	UserId     uint `gorm:"primaryKey;autoIncrement:false; index:index_user_id;comment:订阅者用户id"`

	RemindAt    int  `gorm:"column:remind_at; comment:订阅者个人的提醒时间,见字典项"`
	IsTplRemind bool `gorm:"column:is_tpl_remind; default:0; comment:模板提醒"`
	IsSmsRemind bool `gorm:"column:is_sms_remind; default:0; comment:短信提醒"`
	Status      int  `gorm:"column:status; type:int(2);default:0; comment:订阅状态,0:订阅取消，1:订阅成功"`
}

func (ActivitySubscriber) TableName() string {
	return "wxmp_activity_subscriber"
}

func Init() (err error) {
	err = global.DB.SetupJoinTable(&User{}, "Activities", &Activity{})
	if err != nil {
		global.Log.Error("自定义连接表失败", zap.Any("err", err))
		return err
	}

	err = global.DB.SetupJoinTable(&Activity{}, "Users", &User{})
	if err != nil {
		global.Log.Error("自定义连接表失败", zap.Any("err", err))
		return err
	}
	return nil
}
