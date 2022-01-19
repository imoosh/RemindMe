package wxmp

import (
    "RemindMe/global"
    models "RemindMe/model"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    "errors"
    "github.com/Lofanmi/chinese-calendar-golang/calendar"
    "github.com/golang-module/carbon"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "time"
)

var (
    _pageSize     = 8
    periodicFuncs = map[int]func(t time.Time, n int) time.Time{
        0: func(t time.Time, n int) time.Time { return t },
        1: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddDays(n).Time },
        2: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddWeeks(n).Time },
        3: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddMonthsNoOverflow(n).Time },
        4: func(t time.Time, n int) time.Time { return carbon.Time2Carbon(t).AddMonthsNoOverflow(n).Time },
        5: func(t time.Time, n int) time.Time {
            var tim = t
            var calTmp, calTim *calendar.Calendar
            calTmp = calendar.BySolar(int64(tim.Year()), int64(tim.Month()), int64(tim.Day()), 0, 0, 0)
            calTim = calendar.ByLunar(calTmp.Lunar.GetYear()+int64(n), calTmp.Lunar.GetMonth(), calTmp.Lunar.GetDay(), 0, 0, 0, false)
            tim = time.Date(int(calTim.Solar.GetYear()), time.Month(calTim.Solar.GetMonth()), int(calTim.Solar.GetDay()), 0, 0, 0, 0, time.UTC).Local()
            if tim == time.Unix(0, 0) && calTmp.Lunar.GetDay() == 30 {
                calTim = calendar.ByLunar(calTmp.Lunar.GetYear()+int64(n), calTmp.Lunar.GetMonth(), calTmp.Lunar.GetDay()-1, 0, 0, 0, false)
                tim = time.Date(int(calTim.Solar.GetYear()), time.Month(calTim.Solar.GetMonth()), int(calTim.Solar.GetDay()), 0, 0, 0, 0, time.UTC).Local()
            }
            return time.Date(tim.Year(), tim.Month(), tim.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
        },
    }
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
func (s *ActivityService) ActivityDetail(activityId, subId uint) (activity *wxmp.Activity, err error) {
    if activity, err = s.activityDetail(activityId); err != nil {
        return nil, err
    }
    if activity.Periodic != 0 {
        activity, err = s.getSpecifiedPeriodicSubActivity(activity, subId)
    }
    return
}

// QueryActivities 活动列表
func (s *ActivityService) QueryActivities(userId uint, acType string, cursor string) (list []wxmp.Activity, nextCursor string, err error) {
    var acCursor *activityCursor
    if acCursor, err = parseCursor(cursor); err != nil {
        return
    }

    var activities []wxmp.Activity
    var tim = time.Unix(acCursor.ts, 0)
    if activities, err = s.queryActivities(userId, acType, tim); err != nil {
        return
    }

    list = make([]wxmp.Activity, 0)
    for _, item := range activities {
        if item.Periodic == 0 {
            list = append(list, item)
            continue
        }
        if tmp, err := s.getNotExpiredPeriodicSubActivities(item, acCursor); err == nil {
            list = append(list, tmp...)
        }
    }
    if len(list) == 0 {
        return
    }
    sortActivitiesByTime(list)

    var n = len(list)
    if n >= _pageSize {
        n = _pageSize
        var ac = list[n-1]
        nextCursor = newCursor(ac.Time.Unix(), ac.ID, ac.SubId).String()
    }
    return list[:n], nextCursor, nil
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
func (s *ActivityService) queryActivities(userId uint, typ string, tim time.Time) (list []wxmp.Activity, err error) {
    if tim.Unix() == 0 {
        tim = time.Now()
    }
    // 获取用户所有订阅的活动id
    var activityIds = make([]uint, 0)
    if err = global.DB.Model(&wxmp.ActivitySubscription{}).
        Select("activity_id").
        Where("subscriber_id = ? AND status = 1", userId).
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
        Preload("Subscriptions", func(query *gorm.DB) *gorm.DB { return query.Where("status = 1").Order("updated_at") }).
        Preload("Subscriptions.Subscriber").
        Where("publisher_id = ? OR id IN(?)", userId, activityIds).
        Where("periodic != 0 OR time >= ?", tim.Format(global.MsecLocalTimeFormat))
    if typ != "all" && typ != "" {
        query = query.Where("type = ?", typ)
    }

    if err = query.Order("time").Limit(10).Find(&list).Error; err != nil {
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

// 获取周期性活动的下次未过期的活动信息
func (s *ActivityService) getNotExpiredPeriodicSubActivities(activity wxmp.Activity, cursor *activityCursor) (list []wxmp.Activity, err error) {
    var (
        ac  = activity
        max = _pageSize
        tim = time.Unix(cursor.ts, 0)
    )
    if ac.Periodic > len(periodicFuncs) || ac.Periodic < 0 {
        return nil, errors.New("非法的周期活动")
    }

    // 下一次未过期的周期活动
    list = make([]wxmp.Activity, 0, max)
    for count, i := 0, 0; count < max; i++ {
        t := periodicFuncs[ac.Periodic](activity.Time.Time, i)
        ac.SubId = uint(i)
        ac.Time = models.LocalTime{Time: t}

        if t.Before(tim) || (t.Equal(tim) && ac.ID < cursor.id) || (t.Equal(tim) && ac.ID == cursor.id && ac.SubId <= cursor.subId) {
            continue
        }
        ac.NWeek = int(t.Weekday())
        ac.Lunar = s.jointLunar(t)
        list = append(list, ac)
        count++
    }
    return
}

// 获取周期性活动的指定子活动信息
func (s *ActivityService) getSpecifiedPeriodicSubActivity(old *wxmp.Activity, subId uint) (ac *wxmp.Activity, err error) {
    ac = new(wxmp.Activity)
    *ac = *old

    if old.Periodic > len(periodicFuncs) || old.Periodic < 0 {
        return nil, errors.New("非法的周期活动")
    }
    // 指定未来哪一次的周期活动
    t := periodicFuncs[old.Periodic](old.Time.Time, int(subId))
    if t.Before(time.Now()) {
        return nil, errors.New("活动已结束")
    }
    ac.SubId = subId
    ac.Time = models.LocalTime{Time: t}
    ac.NWeek = int(t.Weekday())
    ac.Lunar = s.jointLunar(t)
    return
}

func (s *ActivityService) jointLunar(t time.Time) string {
    lunarMonth := carbon.Time2Carbon(t).Lunar().ToMonthString()
    lunarDay := carbon.Time2Carbon(t).Lunar().ToDayString()
    if lunarMonth == "十一" {
        lunarMonth = "冬"
    }
    return lunarMonth + "月" + lunarDay
}
