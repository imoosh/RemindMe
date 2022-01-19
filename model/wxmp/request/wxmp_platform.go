package request

// SignatureEventRequest 验证事件调用地址合法
type SignatureEventRequest struct {
    Signature string `json:"signature"`
    Timestamp string `json:"timestamp"`
    Nonce     string `json:"nonce"`
    EchoStr   string `json:"echostr"`
}

type SubscribeMsgEventURLParams struct {
    Signature   string `json:"msg_signature"`
    Timestamp   string `json:"timestamp"`
    Nonce       string `json:"nonce"`
    OpenId      string `json:"openid"`
    EncryptType string `json:"encrypt_type"`
}

type SubscribeMsgEventCommon struct {
    ToUserName   string `json:"ToUserName"`   // 小程序帐号ID
    FromUserName string `json:"FromUserName"` // 用户openid
    CreateTime   string `json:"CreateTime"`   // 时间戳
    MsgType      string `json:"MsgType"`      // 消息类型
    Event        string `json:"Event"`        // 时间类型
    Encrypt      string `json:"Encrypt"`      // 加密数据
}

type SubscribeMsgPopupEvent struct {
    TemplateId            string `json:"TemplateId"`            // 模板id（一次订阅可能有多个id）
    SubscribeStatusString string `json:"SubscribeStatusString"` // 订阅结果（accept接收；reject拒收）
    PopupScene            string `json:"PopupScene"`            // 弹框场景，0代表在小程序页面内
}

type SubscribeMsgChangeEvent struct {
    TemplateId            string `json:"TemplateId"`
    SubscribeStatusString string `json:"SubscribeStatusString"`
}

type SubscribeMsgSentEvent struct {
    TemplateId  string `json:"TemplateId"`  // 模板id（一次订阅可能有多个id）
    MsgID       string `json:"MsgID"`       // 消息id（调用接口时也会返回）
    ErrorCode   string `json:"ErrorCode"`   // 推送结果状态码（0表示成功）
    ErrorStatus string `json:"ErrorStatus"` // 推送结果状态码对应的含义

}

// SubscribeMsgPopupEventRequest 当用户触发订阅消息弹框后
type SubscribeMsgPopupEventRequest struct {
    SubscribeMsgEventCommon
    List SubscribeMsgPopupEvent `json:"List"`
}

// SubscribeMsgChangeEventRequest 当用户在手机端服务通知里消息卡片右上角“...”管理消息时（目前只推送取消订阅的事件，即对消息设置“拒收”）
type SubscribeMsgChangeEventRequest struct {
    SubscribeMsgEventCommon
    List SubscribeMsgChangeEvent `json:"List"`
}

// SubscribeMsgSentEventRequest 调用订阅消息接口发送消息给用户的最终结果
type SubscribeMsgSentEventRequest struct {
    SubscribeMsgEventCommon
    List SubscribeMsgSentEvent `json:"List"`
}
