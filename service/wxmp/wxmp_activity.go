package wxmp

import (
    "RemindMe/global"
    models "RemindMe/model"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    "errors"
    "go.uber.org/zap"
    "gorm.io/gorm/clause"
    "time"
)

type ActivityService struct {
}

// 创建活动
func (s *ActivityService) CreateActivity(userId uint, ac *wxmp.Activity) (err error) {
    var user = wxmp.User{Model: models.Model{ID: userId}}
    if err = global.DB.Debug().Model(&user).Association("Activities").Append(ac); err != nil {
        global.Log.Error("创建活动失败", zap.Any("err", err))
    }
    return
}

// 更新活动
func (s *ActivityService) UpdateActivity(activityId uint, ac *wxmp.Activity) (err error) {
    if err = global.DB.Debug().Where("id = ?", activityId).UpdateColumns(ac).Error; err != nil {
       global.Log.Error("更新活动失败", zap.Any("err", err))
    }
    return
}

// 查询活动详情
func (s *ActivityService) ActivityDetail(activityId uint) (activity *wxmp.Activity, err error) {
    activity = new(wxmp.Activity)
    if err = global.DB.
        Preload("Publisher").
        Preload("Subscriptions").
        Preload("Subscriptions.Subscriber").
        Find(&activity, activityId).Error; err != nil {
        global.Log.Error("获取活动详情失败", zap.Any("err", err))
        return nil, err
    }
    return activity, err
}

// 查询用户参加的所有活动
func (s *ActivityService) QueryActivities(userId uint, typ string) (list []wxmp.Activity, err error) {
    list = make([]wxmp.Activity, 0)
    query := global.DB.Debug().
        Preload("Publisher").
        Preload("Subscriptions").
        Preload("Subscriptions.Subscriber").
        Where("publisher_id = ?", userId)
    if typ != "all" && typ != "" {
        query = query.Where("type = ?", typ)
    }

    if err = query.Find(&list).Error; err != nil {
        global.Log.Error("获取活动列表失败", zap.Any("err", err))
        return nil, err
    }
    global.Log.Error("活动列表", zap.Any("activity", list))
    return list, err
}

// 删除活动
func (s *ActivityService) DeleteActivity(userId, activityId uint) (err error) {
    var activity wxmp.Activity
    if err = global.DB.Where("id = ?", activityId).First(&activity).Error; err != nil {
        global.Log.Error("活动不存在", zap.Any("err", err))
        return err
    }
    // 任务发布者id和用户id必须是同一个人
    if activity.PublisherID != userId {
        return errors.New("无权限删除活动")
    }
    if err = global.DB.Model(&wxmp.Activity{}).Delete("id = ?", activityId).Error; err != nil {
        global.Log.Error("删除活动失败", zap.Any("err", err))
    }
    return
}

// 订阅活动
func (s *ActivityService) SubscribeActivity(userId uint, req *wxmpReq.ActivitySubscribeRequest) (err error) {
    var item = &wxmp.ActivitySubscription{
        SubscriberID: userId,
        ActivityID:   req.Id,
        RemindAt:     req.RemindAt,
        IsTplRemind:  req.IsTplRemind,
        IsSmsRemind:  req.IsSmsRemind,
        Status:       1,
    }
    var updates = map[string]interface{}{
        "updated_at":    time.Now().Format(global.SecLocalTimeFormat),
        "remind_at":     req.RemindAt,
        "is_tpl_remind": req.IsTplRemind,
        "is_sms_remind": req.IsSmsRemind,
        "status":        1,
    }

    if err = global.DB.Debug().Clauses(clause.OnConflict{DoUpdates: clause.Assignments(updates)}).Create(item).Error; err != nil {
        global.Log.Error("新增订阅失败", zap.Any("err", err))
    }
    return
}

// 取消订阅活动
func (s *ActivityService) UnsubscribeActivity(userId, activityId uint) (err error) {
    var item = &wxmp.ActivitySubscription{
        SubscriberID: userId,
        ActivityID:   activityId,
        Status:       0,
    }
    var updates = map[string]interface{}{
        "updated_at": time.Now().Format(global.SecLocalTimeFormat),
        "status":     0,
    }
    if err = global.DB.Debug().Clauses(clause.OnConflict{DoUpdates: clause.Assignments(updates)}).Create(item).Error; err != nil {
        global.Log.Error("取消订阅失败", zap.Any("err", err))
    }
    return
}

// 获取活动所有订阅者
func (s *ActivityService) ActivitySubscribers(activityId uint) (activity *wxmp.Activity, err error) {
    activity = new(wxmp.Activity)
    if err = global.DB.Debug().Preload("Users").Find(&activity, activityId).Error; err != nil {
        global.Log.Error("获取订阅列表失败", zap.Any("err", err))
        return nil, err
    }
    return
}
