package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/wxmp"
    "context"
    "encoding/json"
    "github.com/medivhzhan/weapp/v2"
    "go.uber.org/zap"
    "gorm.io/gorm/clause"
)

type UserService struct {
}

//func (s *UserService) CreateUser(user *wxmp.WxmpUser) {
//    err := global.GVA_DB.Clauses(clause.OnConflict{
//        Columns:   []clause.Column{{Name: "openid"}},
//        DoUpdates: clause.AssignmentColumns([]string{"nickname", "gender", "avatar"}),
//    }).Create(&user).Error
//    if err != nil {
//        global.GVA_LOG.Error("创建小程序用户失败", zap.Any("err", err))
//    }
//}

func (s *UserService) CreateUser(user *wxmp.WxmpUser) {
    if err := global.GVA_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error; err != nil {
        global.GVA_LOG.Error("创建小程序用户失败", zap.Any("err", err))
    }
}

func (s *UserService) GetUserByOpenId(openid string) (u *wxmp.WxmpUser, err error) {
    var user wxmp.WxmpUser
    if err = global.GVA_DB.Where("openid = ?", openid).First(&user).Error; err != nil {
        global.GVA_LOG.Error("查询用户信息失败", zap.Any("openid", openid))
    }
    return &user, err
}

func (s *UserService) UpdateBasicInfo(openid string, user *wxmp.WxmpUser) {
    if err := global.GVA_DB.Where("openid = ?", openid).UpdateColumns(user).Error; err != nil {
        global.GVA_LOG.Error("更新用户信息失败", zap.Any("err", err))
    }
}

func (s *UserService) UpdatePhoneNumber(id uint, phone string) {
    if err := global.GVA_DB.Model(&wxmp.WxmpUser{}).Where("id = ?", id).Update("phone", phone).Error; err != nil {
        global.GVA_LOG.Error("更新手机号号码失败", zap.Any("err", err))
    }
}

// 缓存session_key -> open_id + union_id，等用户登录后再统一入库
func (s *UserService) SetRedisSessionKey(key string, info *weapp.LoginResponse) (err error) {
    var bs []byte
    if bs, err = json.Marshal(info); err != nil {
        global.GVA_LOG.Error("json.Marshal", zap.Any("info", info))
        return err
    }
    // 缓存2小时
    _, err = global.GVA_REDIS.Set(context.Background(), key, string(bs), 7200).Result()
    return err
}

func (s *UserService) GetRedisSessionKey(key string) (*weapp.LoginResponse, error) {
    val, err := global.GVA_REDIS.Get(context.Background(), key).Result()
    if err != nil {
        global.GVA_LOG.Error("获取session_key失败", zap.Any("session_key", key))
        return nil, err
    }

    var info weapp.LoginResponse
    if err = json.Unmarshal([]byte(val), &info); err != nil {
        global.GVA_LOG.Error("json.Unmarshal失败", zap.Any("err", err))
        return nil, err
    }

    // 取到即删除
    global.GVA_REDIS.Del(context.Background(), key)
    return &info, nil
}
