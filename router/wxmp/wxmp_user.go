package wxmp

import (
    v1 "RemindMe/api/v1"
    "RemindMe/middleware"
    "github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
    userRouter := Router.Group("api/v1/user").Use(middleware.OperationRecord())
    var userApi = v1.ApiGroupApp.WxmpApiGroup.UserApi
    {
        userRouter.GET("", userApi.User)                                                  // 用户信息
        userRouter.POST("getWxMiniProgramSessionKey", userApi.GetWxMiniProgramSessionKey) // 获取用户session_key
        userRouter.POST("wxMiniProgramOauth", userApi.WXMiniProgramOauth)                 // 微信小程序登陆
        userRouter.GET("userData", userApi.UserData)                                      // 用户其他信息
        userRouter.POST("accountLogin", userApi.AccountLogin)                             // 用户注册账号
        userRouter.POST("logout", userApi.Logout)                                         // 退出登录
    }
}
