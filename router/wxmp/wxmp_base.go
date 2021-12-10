package wxmp

import (
    v1 "RemindMe/api/v1"
    "github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
    baseRouter := Router.Group("base")
    var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
    {
        baseRouter.POST("login", baseApi.Login)
    }
    return baseRouter
}
