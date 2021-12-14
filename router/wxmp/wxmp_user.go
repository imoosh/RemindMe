package wxmp

import (
    v1 "RemindMe/api/v1"
    "github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
    userRouter := Router.Group("api/v1/user")//.Use(middleware.OperationRecord())
    var userApi = v1.ApiGroupApp.WxmpApiGroup.UserApi
    {
        userRouter.GET("", userApi.UserInfo)                            // 用户信息
        userRouter.GET("userData", userApi.UserData)                    // 用户其他信息
        userRouter.POST("logout", userApi.Logout)                       // 退出登录
        userRouter.POST("wxPhoneNumberAuth", userApi.WXPhoneNumberAuth) //手机号码解密
    }
}
