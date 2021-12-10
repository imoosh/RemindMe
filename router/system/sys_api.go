package system

import (
	v1 "RemindMe/api/v1"
	"RemindMe/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	var apiRouterApi = v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
