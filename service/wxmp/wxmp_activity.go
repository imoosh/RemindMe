package wxmp

import (
	"RemindMe/global"
	models "RemindMe/model"
	"RemindMe/model/wxmp"
	wxmpReq "RemindMe/model/wxmp/request"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"strconv"
)

type ActivityService struct {
}

func (s *ActivityService) CreateActivity(userId uint, ac *wxmp.Activity) (err error) {
	//if err = global.DB.Create(ac).Error; err != nil {
	//    global.Log.Error("创建活动失败", zap.Any("err", err))
	//}
	var user = wxmp.User{Model: models.Model{ID: userId}}
	if err = global.DB.Model(&user).Association("Activities").Append(ac); err != nil {
		global.Log.Error("创建活动失败", zap.Any("err", err))
	}
	return
}

func (s *ActivityService) QueryActivityDetail(activityId string) (activity *wxmp.Activity, err error) {
	activity = new(wxmp.Activity)
	aid, _ := strconv.Atoi(activityId)
	activity.ID = uint(aid)

	if err = global.DB.Preload("User").Preload("Users").Find(&activity).Error; err != nil {
		global.Log.Error("获取订阅列表失败", zap.Any("err", err))
		return nil, err
	}

	return activity, err
}

func (s *ActivityService) QueryActivities(userId uint) ([]wxmp.Activity, error) {
	var err error

	var list = make([]wxmp.Activity, 0)
	if err = global.DB.Preload("User").Where("user_id = ?", userId).Find(&list).Error; err != nil {
		global.Log.Error("查询活动列表失败", zap.Any("err", err))
		return nil, err
	}
	return list, err
}

func (s *ActivityService) DeleteActivity(activityId uint) (err error) {
	if err = global.DB.Model(&wxmp.Activity{}).Delete("id = ?", activityId).Error; err != nil {
		global.Log.Error("删除活动失败", zap.Any("err", err))
	}
	return
}

func (s *ActivityService) SubscribeActivity(userId uint, req *wxmpReq.ActivitySubscribeRequest) (err error) {
	var item = &wxmp.ActivitySubscriber{
		ActivityId:  req.Id,
		UserId:      userId,
		RemindAt:    req.RemindAt,
		IsTplRemind: req.IsTplRemind,
		IsSmsRemind: req.IsSmsRemind,
		Status:      1,
	}

	if err = global.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&item).Error; err != nil {
		global.Log.Error("新增订阅失败", zap.Any("err", err))
	}
	return
}

func (s *ActivityService) UnsubscribeActivity(activityId, userId uint) (err error) {
	if err = global.DB.Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&wxmp.ActivitySubscriber{
			UserId:     userId,
			ActivityId: activityId,
			Status:     0,
		}).Error; err != nil {
		global.Log.Error("取消订阅失败", zap.Any("err", err))
	}
	return
}

func (s *ActivityService) ActivitySubscribers(activityId uint) (activity *wxmp.Activity, err error) {
	activity = new(wxmp.Activity)
	activity.ID = activityId
	if err = global.DB.Preload("Users").Find(&activity).Error; err != nil {
		global.Log.Error("获取订阅列表失败", zap.Any("err", err))
		return nil, err
	}
	global.Log.Error("订阅列表", zap.Any("list", activity))
	return
}
