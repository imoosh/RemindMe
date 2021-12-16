package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/wxmp"
    "go.uber.org/zap"
)

type ActivityService struct {
}

func (s *ActivityService) QueryActivities(id uint) ([]wxmp.WxmpActivity, error) {
    var (
        err  error
        list = make([]wxmp.WxmpActivity, 0)
    )
    if err = global.DB.Model(&wxmp.WxmpActivity{}).Where("publisher_id = ?", id).Find(&list).Error; err != nil {
        global.Log.Error("查询活动列表失败", zap.Any("id", id), zap.Any("err", err))
    }
    return list, err
}

func (s *ActivityService) CreateActivity(ac *wxmp.WxmpActivity) (err error) {
    if err = global.DB.Create(ac).Error; err != nil {
        global.Log.Error("创建活动失败", zap.Any("err", err))
    }
    return
}

func (s *ActivityService) DeleteActivity(id uint) (err error) {
    if err = global.DB.Model(&wxmp.WxmpActivity{}).Delete("id = ?", id).Error; err != nil {
        global.Log.Error("删除活动失败", zap.Any("err", err))
    }
    return
}
