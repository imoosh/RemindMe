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
	"time"
)

type ActivityApi struct {
}

func (api *ActivityApi) ActivityList(c *gin.Context) {
	user := utils.GetWxmpUserInfo(c)
	if user.ID == 0 {
		response.FailWithMessage("获取用户id失败", c)
		return
	}

	list, err := activityService.QueryActivities(user.ID)
	if err != nil {
		response.FailWithMessage("获取活动列表失败", c)
		return
	}

	var acs = make([]wxmpRes.ActivityResponse, 0)
	for _, ac := range list {
		acs = append(acs, wxmpRes.ActivityResponse{
			Id:    ac.ID,
			Title: ac.Title,
			Time: wxmpRes.ActivityTime{
				Solar:    ac.Time.Format("2006-01-02 15:04"),
				IsLunar:  ac.IsLunar,
				Lunar:    ac.Lunar,
				Periodic: ac.Periodic,
				NWeek:    ac.NWeek,
			},
			Addr: wxmpRes.ActivityAddr{
				Address:   ac.Address,
				Latitude:  ac.Latitude,
				Longitude: ac.Longitude,
			},
			RemindAt: ac.RemindAt,
			Remark:   ac.Remark,
			Publisher: wxmpRes.ActivityUser{
				Id:     user.ID,
				Name:   user.Nickname,
				Avatar: user.Avatar,
			},
			Subscribers: nil,
		})
	}

	response.OkWithData(&acs, c)
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
	var ac = wxmp.WxmpActivity{
		PublisherId: id,
		Title:       req.Title,
		Time:        models.LocalTime{Time: t},
		IsLunar:     req.Time.IsLunar,
		Lunar:       req.Time.Lunar,
		Periodic:    req.Time.Periodic,
		NWeek:       req.Time.NWeek,
		Address:     req.Location.Address,
		Latitude:    req.Location.Latitude,
		Longitude:   req.Location.Longitude,
		RemindAt:    req.RemindAt,
		IsTplRemind: req.IsTplRemind,
		IsSmsRemind: req.IsSmsRemind,
		Privacy:     req.Privacy,
		Remark:      req.Remark,
	}

	if err := activityService.CreateActivity(&ac); err != nil {
		response.FailWithMessage("创建活动失败", c)
		return
	}

	response.OkWithData(nil, c)
}

func (api *ActivityApi) UpdateActivity(c *gin.Context) {

}

func (api *ActivityApi) GetActivityDetail(c *gin.Context) {

}

func (api *ActivityApi) DeleteActivity(c *gin.Context) {

}

func (api *ActivityApi) SubscribeActivity(c *gin.Context) {

}
