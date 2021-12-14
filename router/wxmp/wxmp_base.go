package wxmp

import (
    v1 "RemindMe/api/v1"
    "github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
    baseRouter := Router.Group("api/v1")//.Use(middleware.OperationRecord())
    var baseApi = v1.ApiGroupApp.WxmpApiGroup.BaseApi
    var userApi = v1.ApiGroupApp.WxmpApiGroup.UserApi
    {
        baseRouter.GET("index/init", baseApi.IndexInit)                                        // 初始化数据
        baseRouter.POST("user/getWxMiniProgramSessionKey", userApi.GetWxMiniProgramSessionKey) // 获取用户session_key
        baseRouter.POST("user/wxMiniProgramOauth", userApi.WXMiniProgramOauth)                 // 微信小程序登陆
        baseRouter.POST("user/accountLogin", userApi.AccountLogin)                             // 用户注册账号
    }
    return baseRouter
}
