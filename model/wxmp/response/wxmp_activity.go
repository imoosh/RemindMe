package response

// 活动用户信息
type ActivityUser struct {
    Id          uint   `json:"id"`          // 用户id
    Name        string `json:"name"`        // 用户姓名
    Avatar      string `json:"avatar"`      // 用户头像
    Phone       string `json:"phone"`       // 手机号
    IsTplRemind bool   `json:"isTplRemind"` // 模板消息通知
    IsSmsRemind bool   `json:"isSmsRemind"` // 短信通知
}

// 活动时间
type ActivityTime struct {
    Text        string `json:"text"`        // 卡片显示时间
    Solar       string `json:"solar"`       // 日历时间
    IsLunar     bool   `json:"isLunar"`     // 是否为农历
    Lunar       string `json:"lunar"`       // 农历时间
    Periodic    int    `json:"periodic"`    // 周期间隔时间，0/1/7/30/365及355(农历每年)
    NWeek       int    `json:"nWeek"`       // 周期几，1-7
    ObviousDate string `json:"obviousDate"` // 显而易见的日期，如：明天、下周一等
    ObviousTime string `json:"obviousTime"` // 显而易见的时间，如：13:00
}

// 活动地址
type ActivityAddr struct {
    Address   string  `json:"address"`   // 活动地址
    Latitude  float64 `json:"latitude"`  // 地址纬度
    Longitude float64 `json:"longitude"` // 地址精度
}

// 活动信息
type ActivityResponse struct {
    Id           uint           `json:"id"`           // 活动id
    Title        string         `json:"title"`        // 活动名称
    Time         ActivityTime   `json:"time"`         // 活动时间
    Addr         ActivityAddr   `json:"addr"`         // 活动地址
    RemindAt     int            `json:"remindAt"`     // 提醒时间，未订阅前使用创建时设定的默认时间，订阅后可修改
    Remark       string         `json:"remark"`       // 活动备注
    Privacy      int            `json:"privacy"`      // 隐私级别，见字典项
    Publisher    ActivityUser   `json:"publisher"`    // 发布者信息
    Subscribers  []ActivityUser `json:"subscribers"`  // 订阅者信息
    IsSubscribed bool           `json:"isSubscribed"` // 当前用户是否已订阅
    IsPublisher  bool           `json:"isPublisher"`  // 当前用户是否是发布者
}
