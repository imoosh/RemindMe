package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/wxmp"
    "go.uber.org/zap"
    "gorm.io/gorm/clause"
)

type UserService struct {
}

//func (s *UserService) CreateUser(user *wxmp.User) {
//    err := global.DB.Clauses(clause.OnConflict{
//        Columns:   []clause.Column{{Name: "openid"}},
//        DoUpdates: clause.AssignmentColumns([]string{"nickname", "gender", "avatar"}),
//    }).Create(&user).Error
//    if err != nil {
//        global.Log.Error("创建小程序用户失败", zap.Any("err", err))
//    }
//}

func (s *UserService) CreateUser(user *wxmp.User) {
    if err := global.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error; err != nil {
        global.Log.Error("创建小程序用户失败", zap.Any("err", err))
    }
}

func (s *UserService) GetUserById(id uint) (u *wxmp.User, err error) {
    var user wxmp.User
    if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
        global.Log.Error("查询用户信息失败", zap.Any("id", id))
    }
    return &user, err
}

func (s *UserService) GetUserByOpenId(openid string) (u *wxmp.User, err error) {
    var user wxmp.User
    if err = global.DB.Where("openid = ?", openid).First(&user).Error; err != nil {
        global.Log.Error("查询用户信息失败", zap.Any("openid", openid))
    }
    return &user, err
}

func (s *UserService) UpdateBasicInfo(openid string, user *wxmp.User) {
    if err := global.DB.Where("openid = ?", openid).UpdateColumns(user).Error; err != nil {
        global.Log.Error("更新用户信息失败", zap.Any("err", err))
    }
}

func (s *UserService) UpdatePhoneNumber(id uint, phone string) {
    if err := global.DB.Model(&wxmp.User{}).Where("id = ?", id).Update("phone", phone).Error; err != nil {
        global.Log.Error("更新手机号号码失败", zap.Any("err", err))
    }
}
