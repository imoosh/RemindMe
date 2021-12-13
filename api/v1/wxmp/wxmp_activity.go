package wxmp

import (
	"RemindMe/model/common/response"
	wxmpReq "RemindMe/model/wxmp/request"
	wxmpRes "RemindMe/model/wxmp/response"
	"fmt"
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
	fmt.Println(req)
	response.OkWithData(nil, c)
}

func (api *ActivityApi) GetActivityDetail(c *gin.Context) {

}

func (api *ActivityApi) DeleteActivity(c *gin.Context) {

}
