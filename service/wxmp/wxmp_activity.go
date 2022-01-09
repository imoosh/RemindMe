package wxmp

import (
    "RemindMe/global"
    models "RemindMe/model"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    "errors"
    "github.com/golang-module/carbon"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "sort"
    "time"
)

type ActivityService struct {
}

// CreateActivity 创建活动
func (s *ActivityService) CreateActivity(userId uint, ac *wxmp.Activity) (err error) {
    var user = wxmp.User{Model: models.Model{ID: userId}}
    if err = global.DB.Debug().Model(&user).Association("Activities").Append(ac); err != nil {
        global.Log.Error("创建活动失败", zap.Any("err", err))
    }
    return
}

// UpdateActivity 更新活动
func (s *ActivityService) UpdateActivity(activityId uint, ac *wxmp.Activity) (err error) {
    if err = global.DB.Debug().Where("id = ?", activityId).UpdateColumns(ac).Error; err != nil {
        global.Log.Error("更新活动失败", zap.Any("err", err))
    }
    return
}

// ActivityDetail 活动详情
func (s *ActivityService) ActivityDetail(activityId uint) (activity *wxmp.Activity, err error) {
    if activity, err = s.activityDetail(activityId); err != nil {
        return nil, err
    } else if activity.Periodic == 0 {
        return activity, nil
    } else {
        s.calcNextPeriodicActivity(activity)
    }
    return
}

// QueryActivities 活动列表
func (s *ActivityService) QueryActivities(userId uint, typ string) (list []wxmp.Activity, err error) {
    var as []wxmp.Activity

    if as, err = s.queryActivities(userId, typ); err != nil {
        return nil, err
    }
    list = make([]wxmp.Activity, 0, len(as))
    for _, item := range as {
        if item.Periodic != 0 {
            s.calcNextPeriodicActivity(&item)
        }
        list = append(list, item)
    }
    s.sortActivitiesByTime(list)
    return list, nil
}

// activityDetail 查询活动详情
func (s *ActivityService) activityDetail(activityId uint) (activity *wxmp.Activity, err error) {
    if activityId <= 0 {
        return nil, errors.New("活动不存在")
    }
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

// QueryActivities 查询用户参加的所有活动
func (s *ActivityService) queryActivities(userId uint, typ string) (list []wxmp.Activity, err error) {
    // 获取用户所有订阅的活动id
    var activityIds = make([]uint, 0)
    if err = global.DB.Model(&wxmp.ActivitySubscription{}).
        Select("activity_id").
        Where("subscriber_id = ? AND Status = 1", userId).
        Find(&activityIds).Error; err != nil {
        global.Log.Error("获取用户订阅的活动列表失败", zap.Any("userId", userId), zap.Any("err", err))
        return nil, err
    }

    // 获取用户所有创建及订阅的活动信息
    // 1.当前创建或订阅的活动
    // 2.周期性活动或非周期未超时的活动
    // 3.按活动时间排序
    // 4.订阅者按时间排序
    list = make([]wxmp.Activity, 0)
    query := global.DB.Debug().
        Preload("Publisher").
        Preload("Subscriptions", func(query *gorm.DB) *gorm.DB { return query.Order("updated_at") }).
        Preload("Subscriptions.Subscriber").
        Where("publisher_id = ? OR id IN(?)", userId, activityIds).
        Where("periodic != 0 OR time >= ?", time.Now().Format(global.MsecLocalTimeFormat))
    if typ != "all" && typ != "" {
        query = query.Where("type = ?", typ)
    }

    if err = query.Order("time").Find(&list).Error; err != nil {
        global.Log.Error("获取活动列表失败", zap.Any("err", err))
        return nil, err
    }
    return list, err
}

// DeleteActivity 删除活动
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

// SubscribeActivity 订阅活动
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

// UnsubscribeActivity 取消订阅活动
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

// ActivitySubscribers 获取活动所有订阅者
func (s *ActivityService) ActivitySubscribers(activityId uint) (activity *wxmp.Activity, err error) {
    activity = new(wxmp.Activity)
    if err = global.DB.Debug().Preload("Users").Find(&activity, activityId).Error; err != nil {
        global.Log.Error("获取订阅列表失败", zap.Any("err", err))
        return nil, err
    }
    return
}

// Bucket 定义一个通用的结构体
type Bucket struct {
    Slice []wxmp.Activity             //承载以任意结构体为元素构成的Slice
    By    func(a, b interface{}) bool //排序规则函数,当需要对新的结构体slice进行排序时，只需定义这个函数即可
}

func (this Bucket) Len() int { return len(this.Slice) }

func (this Bucket) Swap(i, j int) { this.Slice[i], this.Slice[j] = this.Slice[j], this.Slice[i] }

func (this Bucket) Less(i, j int) bool { return this.By(this.Slice[i], this.Slice[j]) }

func (s *ActivityService) sortActivitiesByTime(list []wxmp.Activity) {
    f := func(a, b interface{}) bool {
        return a.(wxmp.Activity).Time.Time.Before(b.(wxmp.Activity).Time.Time)
    }
    results := Bucket{By: f, Slice: list}
    sort.Sort(results)
}

// 获取指定周期性活动的子活动
func (s *ActivityService) calcNextPeriodicActivity(ac *wxmp.Activity) {
    var funcs = map[int]func(t time.Time, n int) time.Time{
        0: func(t time.Time, n int) time.Time { return t },
        1: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddDays(n).Time },
        2: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddWeeks(n).Time },
        3: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddMonthsNoOverflow(n).Time },
        4: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddMonthsNoOverflow(n).Time },
        5: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddMonthsNoOverflow(n).Time },
    }
    if ac.Periodic > len(funcs) || ac.Periodic < 0 {
        return
    }

    now := time.Now()
    for i := 0; ; i++ {
        t := funcs[ac.Periodic](ac.Time.Time, i)
        if t.Before(now) {
            continue
        }
        ac.Time = models.LocalTime{Time: t}
        ac.NWeek = int(t.Weekday())
        ac.Lunar = s.jointLunar(t)
        break
    }
}

func (s *ActivityService) jointLunar(t time.Time) string {
    lunarMonth := carbon.Time2Carbon(t).Lunar().ToMonthString()
    lunarDay := carbon.Time2Carbon(t).Lunar().ToDayString()
    if lunarMonth == "十一" {
        lunarMonth = "冬"
    }
    return lunarMonth + "月" + lunarDay
}
