package wxmp

import (
    wxmpReq "RemindMe/model/wxmp/request"
)

type PlatformService struct {
}

// HandleSubscribeMsgPopupEvent 当用户触发订阅消息弹框后
func (s *PlatformService) HandleSubscribeMsgPopupEvent(args *wxmpReq.SubscribeMsgPopupEventRequest) {
    //var list []string
    //for _, item := range args.List {
    //    list = append(list, item.)
    //}
    //err = global.Redis.HSet(context.Background(), args.List[0].TemplateId, args.FromUserName, args.).Err()
}

// HandleSubscribeMsgChangeEvent 当用户在手机端服务通知里消息卡片右上角“...”管理消息时（目前只推送取消订阅的事件，即对消息设置“拒收”）
func (s *PlatformService) HandleSubscribeMsgChangeEvent(args *wxmpReq.SubscribeMsgChangeEventRequest) {
}

// HandleSubscribeMsgSentEvent 调用订阅消息接口发送消息给用户的最终结果
func (s *PlatformService) HandleSubscribeMsgSentEvent(args *wxmpReq.SubscribeMsgSentEventRequest) {
}
