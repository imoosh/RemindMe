package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/wxmp"
    "go.uber.org/zap"
)

type ActivityService struct {
}

func (s *ActivityService) CreateActivity(ac *wxmp.WxmpActivity) (err error) {
    if err = global.GVA_DB.Create(ac).Error; err != nil {
        global.GVA_LOG.Error("创建活动失败", zap.Any("err", err))
    }
    return
}
