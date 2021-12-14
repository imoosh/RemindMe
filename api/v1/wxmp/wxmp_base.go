package wxmp

import (
    "RemindMe/model/common/response"
    wxmpRes "RemindMe/model/wxmp/response"
    "github.com/gin-gonic/gin"
)

type BaseApi struct {
}

var (
    activityPeriodicList = []string{
        /* 0:   */ "不重复",
        /* 1:   */ "每天",
        /* 7:   */ "每周",
        /* 30:  */ "每月",
        /* 365: */ "每年",
        /* 355: */ "农历每年",
    }
    activityRemindAtList = []string{
        /* -1:    */ "不提醒",
        /* 0:     */ "准时",
        /* 5:     */ "5分钟前",
        /* 10:    */ "10分钟前",
        /* 15:    */ "15分钟前",
        /* 30:    */ "30分钟前",
        /* 60:    */ "1小时前",
        /* 120:   */ "2小时前",
        /* 1440:  */ "1天前",
        /* 2880:  */ "2天前",
        /* 10080: */ "1周前",
    }
    ActivityPrivacyList = []string{
        /* 0: */ "仅自己可见",
        /* 1: */ "可分享他人",
        /* 2: */ "完全公开",
    }
)

// @Tags Base
// @Summary 初始化数据
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *BaseApi) IndexInit(c *gin.Context) {

    var res = wxmpRes.BaseInitResponse{
        ActivityPeriodicList: activityPeriodicList,
        ActivityRemindAtList: activityRemindAtList,
        ActivityPrivacyList:  ActivityPrivacyList,
    }

    response.OkWithData(&res, c)
}
