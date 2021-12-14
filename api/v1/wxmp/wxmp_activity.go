package wxmp

import (
    "RemindMe/model/common/response"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    wxmpRes "RemindMe/model/wxmp/response"
    "RemindMe/utils"
    "github.com/gin-gonic/gin"
)

type ActivityApi struct {
}

func (api *ActivityApi) GetActivityList(c *gin.Context) {
    //var l wxmpReq.ActivityListRequest
    //_ = c.ShouldBindJSON(&l)

    var res wxmpRes.ActivityListResponse
    response.OkWithData(&res, c)
}

func (api *ActivityApi) CreateActivity(c *gin.Context) {
    var req wxmpReq.ActivityCreateRequest
    _ = c.ShouldBindJSON(&req)

    id := utils.GetWxmpUserID(c)
    if id == 0 {
        response.FailWithMessage("创建活动失败", c)
        return
    }

    var ac = wxmp.WxmpActivity{
        PublisherId: id,
        Title:       req.Title,
        Time:        req.Time.Solar,
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

func (api *ActivityApi) GetActivityDetail(c *gin.Context) {

}

func (api *ActivityApi) DeleteActivity(c *gin.Context) {

}
