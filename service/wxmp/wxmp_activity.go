package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    "go.uber.org/zap"
    "gorm.io/gorm/clause"
)

type ActivityService struct {
}

func (s *ActivityService) QueryActivity(id string) (*wxmp.WxmpActivity, error) {
    var (
        err error
        ac  = new(wxmp.WxmpActivity)
    )
    if err = global.DB.Model(&wxmp.WxmpActivity{}).Where("id = ?", id).Find(ac).Error; err != nil {
        global.Log.Error("查询活动列表失败", zap.Any("id", id), zap.Any("err", err))
    }
    return ac, err
}

func (s *ActivityService) QueryActivities(id uint) ([]wxmp.WxmpActivity, error) {
    var (
        err  error
        list = make([]wxmp.WxmpActivity, 0)
    )
    if err = global.DB.Model(&wxmp.WxmpActivity{}).Where("publisher_id = ?", id).Find(&list).Error; err != nil {
        global.Log.Error("查询活动列表失败", zap.Any("id", id), zap.Any("err", err))
    }
    return list, err
}

func (s *ActivityService) CreateActivity(ac *wxmp.WxmpActivity) (err error) {
    if err = global.DB.Create(ac).Error; err != nil {
        global.Log.Error("创建活动失败", zap.Any("err", err))
    }
    return
}

func (s *ActivityService) DeleteActivity(id uint) (err error) {
    if err = global.DB.Model(&wxmp.WxmpActivity{}).Delete("id = ?", id).Error; err != nil {
        global.Log.Error("删除活动失败", zap.Any("err", err))
    }
    return
}

func (s *ActivityService) SubscribeActivity(userId uint, req *wxmpReq.ActivitySubscribeRequest) (err error) {
    var item = &wxmp.WxmpActivitySubscribers{
        ActivityId:       req.Id,
        SubscriberUserId: userId,
        RemindAt:         req.RemindAt,
        IsTplRemind:      req.IsTplRemind,
        IsSmsRemind:      req.IsSmsRemind,
        Status:           1,
    }

    if err = global.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&item).Error; err != nil {
        global.Log.Error("新增订阅失败", zap.Any("err", err))
    }
    return
}

func (s *ActivityService) UnsubscribeActivity(activityId, userId uint) (err error) {
    if err = global.DB.Create(&wxmp.WxmpActivitySubscribers{
        SubscriberUserId: userId,
        ActivityId:       activityId,
        Status:           0,
    }).Error; err != nil {
        global.Log.Error("取消订阅失败", zap.Any("err", err))
    }
    return
}

func (s *ActivityService) ActivitySubscribers(activityId uint) (err error) {

    query := global.DB.Model(&wxmp.WxmpActivitySubscribers{}).Where("activity_id = ?", activityId).Order("updatedAt")
    query = query.Joins("LEFT JOIN (?) ON")
    return
}