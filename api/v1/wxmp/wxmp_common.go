package wxmp

import (
    "RemindMe/model/common/response"
    wxmpReq "RemindMe/model/wxmp/request"
    "fmt"
    "github.com/gin-gonic/gin"
)

type CommonApi struct {
}

// @Tags Base
// @Summary 初始化数据
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *CommonApi) IndexInit(c *gin.Context) {
    var l wxmpReq.Login
    _ = c.ShouldBindJSON(&l)

    fmt.Println(l)
    response.OkWithMessage("修改成功", c)
}
