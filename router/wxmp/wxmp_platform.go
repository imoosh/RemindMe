package wxmp

import (
    v1 "RemindMe/api/v1"
    "github.com/gin-gonic/gin"
)

type PlatformRouter struct {
}

func (r *PlatformRouter) InitPlatformRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
    baseRouter := Router.Group("api/v1") //.Use(middleware.OperationRecord())
    var platformApi = v1.ApiGroupApp.WxmpApiGroup.PlatformApi
    {
        baseRouter.GET("wechat/notification", platformApi.CheckSignatureEvent) // 检验事件通知地址是否合法
        baseRouter.POST("wechat/notification", platformApi.SubscribeMsgEvent)  // 订阅模板消息事件通知
    }
    return baseRouter
}
