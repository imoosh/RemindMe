package wxmp

import (
    v1 "RemindMe/api/v1"
    "RemindMe/middleware"
    "github.com/gin-gonic/gin"
)

type CommonRouter struct {
}

func (s *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
    userRouter := Router.Group("api/v1/index/").Use(middleware.OperationRecord())
    var userApi = v1.ApiGroupApp.WxmpApiGroup.CommonApi
    {
        userRouter.GET("init", userApi.IndexInit) // 初始化数据
    }
}
