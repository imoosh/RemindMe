package utils

import (
    "RemindMe/global"
    wxmpReq "RemindMe/model/wxmp/request"
    "github.com/gin-gonic/gin"
)

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetWxmpUserID(c *gin.Context) uint {
    if claims, exists := c.Get("claims"); !exists {
        global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
        return 0
    } else {
        waitUse := claims.(*wxmpReq.CustomClaims)
        return waitUse.ID
    }
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetWxmpUserInfo(c *gin.Context) *wxmpReq.CustomClaims {
    if claims, exists := c.Get("claims"); !exists {
        global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件!")
        return nil
    } else {
        waitUse := claims.(*wxmpReq.CustomClaims)
        return waitUse
    }
}
