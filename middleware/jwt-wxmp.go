package middleware

import (
    "RemindMe/model/common/response"
    "RemindMe/service"
    "RemindMe/utils"
    "github.com/gin-gonic/gin"
)

var jwtWxmpService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTWXMPAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
        token := c.Request.Header.Get("token")
        if token == "" {
            response.FailWithMessage("未登录或非法访问", c)
            c.Abort()
            return
        }
        if jwtService.IsBlacklist(token) {
            response.FailWithMessage("您的帐户异地登陆或令牌失效", c)
            c.Abort()
            return
        }
        j := utils.NewWxmpJWT()
        // parseToken 解析token包含的信息
        claims, err := j.ParseToken(token)
        if err != nil {
            if err == utils.TokenExpired {
                response.TokenExpired(c)
                c.Abort()
                return
            }
            response.FailWithMessage(err.Error(), c)
            c.Abort()
            return
        }
        c.Set("claims", claims)
        c.Next()
    }
}
