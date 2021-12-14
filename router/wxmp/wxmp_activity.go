package wxmp

import (
	v1 "RemindMe/api/v1"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct {
}

func (s *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	activityRouter := Router.Group("api/v1/activity") //.Use(middleware.OperationRecord())
	var activityApi = v1.ApiGroupApp.WxmpApiGroup.ActivityApi
	{
		activityRouter.GET("list", activityApi.ActivityList)
		activityRouter.POST("create", activityApi.CreateActivity)
		activityRouter.POST("update", activityApi.UpdateActivity)
		activityRouter.POST("delete", activityApi.DeleteActivity)
		activityRouter.POST("subscribe", activityApi.SubscribeActivity)
	}
	return activityRouter
}
