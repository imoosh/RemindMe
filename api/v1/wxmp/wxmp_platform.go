package wxmp

import (
    "RemindMe/global"
    wxmpReq "RemindMe/model/wxmp/request"
    "RemindMe/utils"
    "encoding/json"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "net/http"
)

type PlatformApi struct {
}

// CheckSignatureEvent 验证事件调用地址合法
func (api *PlatformApi) CheckSignatureEvent(c *gin.Context) {
    var args wxmpReq.SignatureEventRequest
    args.Signature = c.Query("signature")
    args.Timestamp = c.Query("timestamp")
    args.Nonce = c.Query("nonce")
    args.EchoStr = c.Query("echostr")

    if !utils.CheckSignature(args.Signature, args.Timestamp, args.Nonce, "") {
        global.Log.Error("签名校验失败")
    }
    c.String(http.StatusOK, args.EchoStr)
}

// SubscribeMsgEvent 订阅消息通知事件
func (api *PlatformApi) SubscribeMsgEvent(c *gin.Context) {
    var args wxmpReq.SubscribeMsgEventURLParams
    args.Signature = c.Query("msg_signature")
    args.Timestamp = c.Query("timestamp")
    args.EncryptType = c.Query("encrypt_type")
    args.OpenId = c.Query("openid")
    args.Nonce = c.Query("nonce")

    var header wxmpReq.SubscribeMsgEventCommon
    _ = c.ShouldBindJSON(&header)
    str, _, err := utils.DecryptWeComEventMsg(args.Signature, args.Timestamp, args.Nonce, header.Encrypt)
    if err != nil {
        global.Log.Error("解码数据失败", zap.Any("err", err))
    }

    switch header.Event {
    case "subscribe_msg_popup_event":
        var args wxmpReq.SubscribeMsgPopupEventRequest
        _ = json.Unmarshal([]byte(str), &args)
        // 当用户触发订阅消息弹框后
        platformService.HandleSubscribeMsgPopupEvent(&args)
    case "subscribe_msg_change_event":
        var args wxmpReq.SubscribeMsgChangeEventRequest
        _ = json.Unmarshal([]byte(str), &args)
        // 当用户在手机端服务通知里消息卡片右上角“...”管理消息时（目前只推送取消订阅的事件，即对消息设置“拒收”）
        platformService.HandleSubscribeMsgChangeEvent(&args)
    case "subscribe_msg_sent_event":
        var args wxmpReq.SubscribeMsgSentEventRequest
        _ = json.Unmarshal([]byte(str), &args)
        // 调用订阅消息接口发送消息给用户的最终结果
        platformService.HandleSubscribeMsgSentEvent(&args)
    }

    c.String(http.StatusOK, "success")
}
