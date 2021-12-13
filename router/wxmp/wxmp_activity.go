package wxmp

import (
	v1 "RemindMe/api/v1"
	"RemindMe/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct {
}

func (s *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	activityRouter := Router.Group("api/v1/activity").Use(middleware.OperationRecord())
	var activityApi = v1.ApiGroupApp.WxmpApiGroup.ActivityApi
	{
		activityRouter.POST("create", activityApi.CreateActivity)
	}
	return activityRouter
}
