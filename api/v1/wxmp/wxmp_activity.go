package wxmp

import (
    "RemindMe/global"
    models "RemindMe/model"
    "RemindMe/model/common/response"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    wxmpRes "RemindMe/model/wxmp/response"
    "RemindMe/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "strconv"
    "time"
)

type ActivityApi struct {
}

func (api *ActivityApi) CreateActivity(c *gin.Context) {
    var req wxmpReq.ActivityCreateRequest
    _ = c.ShouldBindJSON(&req)

    id := utils.GetWxmpUserID(c)
    if id == 0 {
        response.FailWithMessage("创建活动失败", c)
        return
    }

    t, err := time.Parse(string("2006-01-02 15:04"), req.Time.Solar)
    if err != nil {
        global.Log.Error("解析时间失败", zap.Any("err", err))
        response.FailWithMessage("创建活动失败", c)
        return
    }
    var ac = wxmp.Activity{
        Title:     req.Title,
        Time:      models.LocalTime{Time: t},
        IsLunar:   req.Time.IsLunar,
        Lunar:     req.Time.Lunar,
        Periodic:  req.Time.Periodic,
        NWeek:     req.Time.NWeek,
        Address:   req.Location.Address,
        Latitude:  req.Location.Latitude,
        Longitude: req.Location.Longitude,
        RemindAt:  req.RemindAt,
        Privacy:   req.Privacy,
        Remark:    req.Remark,
    }

    if err := activityService.CreateActivity(id, &ac); err != nil {
        response.FailWithMessage("创建活动失败", c)
        return
    }

    response.OkWithData(nil, c)
}

func (api *ActivityApi) ActivityList(c *gin.Context) {
    user := utils.GetWxmpUserInfo(c)
    if user.ID == 0 {
        response.FailWithMessage("获取用户id失败", c)
        return
    }

    // 通过用户id获取所有相关的活动
    activities, err := activityService.QueryActivities(user.ID)
    if err != nil {
        response.FailWithMessage("获取活动列表失败", c)
        return
    }

    var list = make([]wxmpRes.ActivityResponse, 0)
    for _, ac := range activities {
        var timeText string
        if !ac.IsLunar {
            timeText = ac.Time.Format("2006-01-02 15:04") + " " + getWeekdayString(ac.NWeek)
        } else {

        }

        res := wxmpRes.ActivityResponse{
            Id:    ac.ID,
            Title: ac.Title,
            Time: wxmpRes.ActivityTime{
                Text:        timeText,
                Solar:       ac.Time.Format("2006-01-02 15:04"),
                IsLunar:     ac.IsLunar,
                Lunar:       ac.Lunar,
                Periodic:    ac.Periodic,
                NWeek:       ac.NWeek,
                ObviousDate: getObviousDate(ac.Time.Time),
                ObviousTime: getObviousTime(ac.Time.Time),
            },
            Addr: wxmpRes.ActivityAddr{
                Address:   ac.Address,
                Latitude:  ac.Latitude,
                Longitude: ac.Longitude,
            },
            Publisher: wxmpRes.ActivityUser{
                Id:     ac.Publisher.ID,
                Name:   ac.Publisher.Nickname,
                Avatar: ac.Publisher.Avatar,
            },
            SubscriberIndex: -1,
        }

        // 当前用户是否为发布者
        if res.Publisher.Id == user.ID {
            res.IsPublisher = true
        }
        // 当前用户是否订阅此活动
        for idx, item := range ac.Subscriptions {
            res.Subscribers = append(res.Subscribers, wxmpRes.ActivityUser{
                Id:     item.ID,
                Name:   item.Subscriber.Nickname,
                Avatar: item.Subscriber.Avatar,
            })
            if item.ID == user.ID {
                res.IsSubscribed = true
                res.SubscriberIndex = idx
                break
            }
        }
        list = append(list, res)
    }

    response.OkWithData(&list, c)
}

func (api *ActivityApi) UpdateActivity(c *gin.Context) {

}

func (api *ActivityApi) ActivityDetail(c *gin.Context) {
    activityId, err := strconv.Atoi(c.Query("id"))
    if err != nil {
        response.FailWithMessage("活动id不存在", c)
        return
    }
    user := utils.GetWxmpUserInfo(c)
    if user.ID == 0 {
        response.FailWithMessage("获取用户id失败", c)
        return
    }

    // 查询活动信息
    ac, err := activityService.ActivityDetail(uint(activityId))
    if err != nil {
        response.FailWithMessage("活动不存在", c)
        return
    }
    global.Log.Error("活动详情", zap.Any("activity", ac))

    var timeText string
    if !ac.IsLunar {
        timeText = ac.Time.Format("2006-01-02 15:04") + " " + getWeekdayString(ac.NWeek)
    } else {

    }

    res := wxmpRes.ActivityResponse{
        Id:    ac.ID,
        Title: ac.Title,
        Time: wxmpRes.ActivityTime{
            Text:        timeText,
            Solar:       ac.Time.Format("2006-01-02 15:04"),
            IsLunar:     ac.IsLunar,
            Lunar:       ac.Lunar,
            Periodic:    ac.Periodic,
            NWeek:       ac.NWeek,
            ObviousDate: getObviousDate(ac.Time.Time),
            ObviousTime: getObviousTime(ac.Time.Time),
        },
        Addr: wxmpRes.ActivityAddr{
            Address:   ac.Address,
            Latitude:  ac.Latitude,
            Longitude: ac.Longitude,
        },
        RemindAt:        ac.RemindAt,
        Remark:          ac.Remark,
        Privacy:         ac.Privacy,
        IsPublisher:     false,
        IsSubscribed:    false,
        SubscriberIndex: -1,
        Publisher: wxmpRes.ActivityUser{
            Id:     ac.Publisher.ID,
            Name:   ac.Publisher.Nickname,
            Avatar: ac.Publisher.Avatar,
        },
        Subscribers: nil,
    }
    // 当前用户是否为发布者
    if res.Publisher.Id == user.ID {
        res.IsPublisher = true
    }
    var subscriber []wxmpRes.ActivityUser
    for idx, item := range ac.Subscriptions {
        if item.Status == 0 {
            continue
        }
        subscriber = append(subscriber, wxmpRes.ActivityUser{
            Id:          item.Subscriber.ID,
            Name:        item.Subscriber.Nickname,
            Avatar:      item.Subscriber.Avatar,
            Phone:       item.Subscriber.Phone,
            IsTplRemind: item.IsTplRemind,
            IsSmsRemind: item.IsSmsRemind,
        })
        // 订阅者id与用户id相同，表示当前用户已订阅活动
        if item.Subscriber.ID == user.ID {
            res.IsSubscribed = true
            res.SubscriberIndex = idx
        }
    }
    res.Subscribers = subscriber

    response.OkWithData(&res, c)
}

func (api *ActivityApi) DeleteActivity(c *gin.Context) {
    var req wxmpReq.ActivityDeleteRequest
    _ = c.ShouldBindJSON(&req)
    id := utils.GetWxmpUserID(c)

    if err := activityService.DeleteActivity(id, req.Id); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    response.Ok(c)
}

func (api *ActivityApi) SubscribeActivity(c *gin.Context) {
    id := utils.GetWxmpUserID(c)

    var req wxmpReq.ActivitySubscribeRequest
    _ = c.ShouldBindJSON(&req)

    if err := activityService.SubscribeActivity(id, &req); err != nil {
        response.FailWithMessage("订阅成功", c)
        return
    }
    response.OkWithMessage("订阅成功", c)
}

func (api *ActivityApi) UnsubscribeActivity(c *gin.Context) {
    var (
        id     int
        err    error
        userId = utils.GetWxmpUserID(c)
    )

    activityId := c.Query("id")
    if id, err = strconv.Atoi(activityId); err != nil {
        response.FailWithMessage("取消订阅活动失败", c)
        return
    }

    if err = activityService.UnsubscribeActivity(userId, uint(id)); err != nil {
        response.FailWithMessage("取消订阅失败", c)
        return
    }
    response.OkWithMessage("取消订阅成功", c)
}

func getObviousDate(tim time.Time) (d string) {
    days := diffDays(time.Now(), tim)
    half := getObviousHalfDate(tim)
    //nweek := tim.Weekday()
    //nextSunday := time.Now()
    //lastMonday := time.Now()
    switch days {
    case -2:
        d = "前天" + half
    case -1:
        d = "昨天" + half
    case 0:
        d = "今天" + half
    case 1:
        d = "明天" + half
    case 2:
        d = "后天" + half
    default:
        d = strconv.Itoa(int(tim.Month())) + "月" + strconv.Itoa(tim.Day()) + "日"
    }
    return d
}

func getObviousTime(tim time.Time) (t string) {
    return tim.Format("15:04")
}

func getObviousHalfDate(t time.Time) string {
    var h, m = t.Hour(), t.Minute()
    if h >= 0 && h < 6 {
        return "凌晨"
    } else if h >= 6 && h < 8 {
        return "早上"
    } else if h >= 8 && h < 11 || (h == 11 && m < 30) {
        return "上午"
    } else if h >= 13 && h < 18 || (h == 12 && m > 30) {
        return "下午"
    } else if h >= 18 && h < 24 {
        return "晚上"
    } else {
        return "中午"
    }
}

func diffDays(t1, t2 time.Time) (days int) {
    var (
        df = t2.Sub(t1)
        du = time.Hour * 24
    )

    // >1天时间间隔
    if df > du || -df > du {
        tmp := int(df / du)
        days += tmp
        t1 = t1.Add(time.Duration(tmp) * du)
    }

    // 不到一天的时间内，如果跨天，days+1
    if t1.Format("20060102") != t2.Format("20060102") {
        if df > 0 {
            days += 1
        } else {
            days -= 1
        }
    }

    return days
}

func getWeekdayString(n int) string {
    var w = map[int]string{
        1: "周一",
        2: "周二",
        3: "周三",
        4: "周四",
        5: "周五",
        6: "周六",
        7: "周日",
    }
    return w[n]
}
