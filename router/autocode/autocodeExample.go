package autocode

import (
	v1 "RemindMe/api/v1"
	"RemindMe/middleware"
	"github.com/gin-gonic/gin"
)

type AutoCodeExampleRouter struct {
}

func (s *AutoCodeExampleRouter) InitSysAutoCodeExampleRouter(Router *gin.RouterGroup) {
	autoCodeExampleRouter := Router.Group("autoCodeExample").Use(middleware.OperationRecord())
	autoCodeExampleRouterWithoutRecord := Router.Group("autoCodeExample")
	var autoCodeExampleApi = v1.ApiGroupApp.AutoCodeApiGroup.AutoCodeExampleApi
	{
		autoCodeExampleRouter.POST("createSysAutoCodeExample", autoCodeExampleApi.CreateAutoCodeExample)   // 新建AutoCodeExample
		autoCodeExampleRouter.DELETE("deleteSysAutoCodeExample", autoCodeExampleApi.DeleteAutoCodeExample) // 删除AutoCodeExample
		autoCodeExampleRouter.PUT("updateSysAutoCodeExample", autoCodeExampleApi.UpdateAutoCodeExample)    // 更新AutoCodeExample
	}
	{
		autoCodeExampleRouterWithoutRecord.GET("findSysAutoCodeExample", autoCodeExampleApi.FindAutoCodeExample)       // 根据ID获取AutoCodeExample
		autoCodeExampleRouterWithoutRecord.GET("getSysAutoCodeExampleList", autoCodeExampleApi.GetAutoCodeExampleList) // 获取AutoCodeExample列表
	}
}
