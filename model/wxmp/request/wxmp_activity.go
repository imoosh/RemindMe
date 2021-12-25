package request

// 活动用户信息
type ActivityUser struct {
	Id     int64  `json:"id"`     // 用户id
	Name   string `json:"name"`   // 用户姓名
	Avatar string `json:"avatar"` // 用户头像
}

// 活动地址
type ActivityLocation struct {
	Address   string  `json:"address"`   // 活动地址
	Latitude  float64 `json:"latitude"`  // 地址纬度
	Longitude float64 `json:"longitude"` // 地址精度
}

// 创建活动请求
type ActivityCreateRequest struct {
	Type     string           `json:"type"`     // 活动类型
	Title    string           `json:"title"`    // 活动名称
	Time     string           `json:"time"`     // 日历时间
	Lunar    string           `json:"lunar"`    // 农历时间
	Periodic int              `json:"periodic"` // 周期间隔时间，0/1/7/30/365及355(农历每年)
	NWeek    int              `json:"nWeek"`    // 周期几，1-7
	Location ActivityLocation `json:"location"` // 活动地址
	Remark   string           `json:"remark"`   // 活动备注
	//Publisher   ActivityUser     `json:"publisher"`   // 发布者信息
	//Subscribers []ActivityUser   `json:"subscribers"` // 订阅者信息
	Privacy     int  `json:"privacy"`     // 隐私级别，见字典项
	RemindAt    int  `json:"remindAt"`    // 创建时设置提醒时间,见字典项"
	IsTplRemind bool `json:"isTplRemind"` // 模板提醒
	IsSmsRemind bool `json:"isSmsRemind"` // 短信提醒
}

// 更新活动请求
type ActivityUpdateRequest struct {
	Id       uint             `json:"id"`       // 活动id
	Type     string           `json:"type"`     // 活动类型
	Title    string           `json:"title"`    // 活动名称
	Time     string           `json:"time"`     // 日历时间
	Lunar    string           `json:"lunar"`    // 农历时间
	Periodic int              `json:"periodic"` // 周期间隔时间，0/1/7/30/365及355(农历每年)
	NWeek    int              `json:"nWeek"`    // 周期几，1-7
	Location ActivityLocation `json:"location"` // 活动地址
	Remark   string           `json:"remark"`   // 活动备注
	//Publisher   ActivityUser     `json:"publisher"`   // 发布者信息
	//Subscribers []ActivityUser   `json:"subscribers"` // 订阅者信息
	Privacy     int  `json:"privacy"`     // 隐私级别，见字典项
	RemindAt    int  `json:"remindAt"`    // 创建时设置提醒时间,见字典项"
	IsTplRemind bool `json:"isTplRemind"` // 模板提醒
	IsSmsRemind bool `json:"isSmsRemind"` // 短信提醒
}

// 删除活动请求
type ActivityDeleteRequest struct {
	Id          uint `json:"id"`
	PublisherId uint `json:"publisherId"`
}

// 活动详情请求
type ActivityDetailRequest struct {
	Id          uint `json:"id"`
	PublisherId uint `json:"publisherId"`
}

// 订阅活动请求
type ActivitySubscribeRequest struct {
	Id          uint `json:"id"`
	RemindAt    int  `json:"remindAt"`
	IsTplRemind bool `json:"isTplRemind"`
	IsSmsRemind bool `json:"isSmsRemind"`
}
