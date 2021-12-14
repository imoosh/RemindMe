package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/common/response"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    wxmpRes "RemindMe/model/wxmp/response"
    "RemindMe/utils"
    "encoding/json"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "github.com/medivhzhan/weapp/v2"
    "github.com/patrickmn/go-cache"
    "go.uber.org/zap"
    "time"
)

const (
    AppId     = "wx5a1de9e90dfc4020"
    AppSecret = "c84f5ecec2fd8fb42e7b993d26793a2e"
)

var (
    tokenCache = cache.New(5*time.Minute, 10*time.Minute)
)

type UserApi struct {
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) AccountLogin(c *gin.Context) {
    var l wxmpReq.Login
    _ = c.ShouldBindJSON(&l)

    fmt.Println(l)
    response.OkWithMessage("修改成功", c)
}

// @Tags Base
// @Summary 用户信息
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) UserInfo(c *gin.Context) {
    info := utils.GetWxmpUserInfo(c)
    if info == nil {
        global.GVA_LOG.Error("获取用户信息失败")
        response.FailWithMessage("获取用户信息失败", c)
        return
    }

    response.OkWithData(&wxmpRes.UserDataResponse{Avatar: info.Avatar, NickName: info.Nickname}, c)
}

// @Tags Base
// @Summary 用户登出
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) Logout(c *gin.Context) {
    var l wxmpReq.Login
    _ = c.ShouldBindJSON(&l)

    fmt.Println(l)
    response.OkWithMessage("修改成功", c)
}

// @Tags Base
// @Summary 获取用户session_key
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) GetWxMiniProgramSessionKey(c *gin.Context) {
    var req wxmpReq.GetWxMiniProgramSessionKeyRequest
    _ = c.ShouldBindJSON(&req)

    info, err := weapp.Login(AppId, AppSecret, req.Code)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }

    // 根据openid+unionid创建小程序用户入库，用户头像性别+手机号码后续更新，并缓存session_key:openid+unionid
    userService.CreateUser(&wxmp.WxmpUser{OpenID: info.OpenID, UnionID: info.UnionID})

    response.OkWithData(&wxmpRes.GetWxMiniProgramSessionKeyResponse{
        SessionKey: info.SessionKey,
        OpenId:     info.OpenID,
    }, c)
}

// @Tags Base
// @Summary 微信小程序登陆
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) WXMiniProgramOauth(c *gin.Context) {
    var req wxmpReq.WXMiniProgramOauthRequest
    _ = c.ShouldBindJSON(&req)

    info, err := weapp.DecryptUserInfo(req.SessionKey, req.RawData, req.EncryptedData, req.Signature, req.IV)
    if err != nil {
        global.GVA_LOG.Error("数据解密失败!", zap.Any("err", err))
        response.FailWithMessage(err.Error(), c)
        return
    }

    // 获取openid
    openid := c.Request.Header.Get("openid")
    if openid == "" {
        global.GVA_LOG.Error("openid不存在")
        response.FailWithMessage("登陆失败", c)
    }

    // 通过openid更新用户基本信息
    var u = wxmp.WxmpUser{Nickname: info.Nickname, Gender: info.Gender, Avatar: info.Avatar}
    userService.UpdateBasicInfo(openid, &u)
    user, err := userService.GetUserByOpenId(openid)
    if err != nil {
        global.GVA_LOG.Error("查询用户信息失败", zap.Any("err", err))
        return
    }

    info.OpenID = user.OpenID
    info.UnionID = user.UnionID
    token, err := b.tokenNext(c, user.ID, info)
    if err != nil {
        global.GVA_LOG.Error("获取token失败!", zap.Any("err", err))
        response.FailWithMessage(err.Error(), c)
        return
    }

    var res = wxmpRes.WXMiniProgramOauthResponse{
        Token: token,
    }
    response.OkWithDetailed(&res, "登录成功", c)
}

// @Tags Base
// @Summary 用户其他信息
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) UserData(c *gin.Context) {

    info := utils.GetWxmpUserInfo(c)
    if info == nil {
        global.GVA_LOG.Error("获取用户信息失败")
        response.Fail(c)
        return
    }

    var res = wxmpRes.UserDataResponse{
        Avatar:   info.Avatar,
        NickName: info.Nickname,
    }
    response.OkWithData(&res, c)
}

// 登录以后签发jwt
func (b *UserApi) tokenNext(c *gin.Context, id uint, user *weapp.UserInfo) (token string, err error) {
    j := &utils.WxmpJWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
    claims := wxmpReq.CustomClaims{
        ID:         id,
        OpenID:     user.OpenID,
        Nickname:   user.Nickname,
        Gender:     user.Gender,
        Province:   user.Province,
        Language:   user.Language,
        Country:    user.Country,
        City:       user.City,
        Avatar:     user.Avatar,
        UnionID:    user.UnionID,
        BufferTime: 3600, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
        StandardClaims: jwt.StandardClaims{
            NotBefore: time.Now().Unix() - 1000, // 签名生效时间
            //ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
            ExpiresAt: time.Now().Unix() + 30*24*3600, // 过期时间 7天  配置文件
            Issuer:    "7gcat",                        // 签名的发行者
        },
    }
    token, err = j.CreateToken(claims)
    if err != nil {
        global.GVA_LOG.Error("获取token失败!", zap.Any("err", err))
        response.FailWithMessage("获取token失败", c)
        return "", err
    }
    //return token, nil
    var userbs []byte
    userbs, err = json.Marshal(user)
    if err != nil {
        global.GVA_LOG.Error("json.Marshal失败!", zap.Any("err", err))
        return
    }
    userStr := string(userbs)

    if err, _ = jwtService.GetRedisJWT(token); err == redis.Nil {
        if err = jwtService.SetRedisJWT(userStr, token); err != nil {
            global.GVA_LOG.Error("设置登录状态失败!", zap.Any("err", err))
            return
        }
    } else if err != nil {
        global.GVA_LOG.Error("设置登录状态失败!", zap.Any("err", err))
        return
    } else {
        if err = jwtService.SetRedisJWT(token, userStr); err != nil {
            return
        }
    }

    return token, nil
}

// @Tags Base
// @Summary 微信小程序手机号码授权
// @Produce  application/json
// @Param data body wxmpReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *UserApi) WXPhoneNumberAuth(c *gin.Context) {

    info := utils.GetWxmpUserInfo(c)
    if info == nil {
        global.GVA_LOG.Error("获取用户信息失败")
        response.Fail(c)
        return
    }

    var req wxmpReq.WXPhoneNumberAuthRequest
    _ = c.ShouldBindJSON(&req)

    mobile, err := weapp.DecryptMobile(req.SessionKey, req.EncryptedData, req.IV)
    if err != nil {
        global.GVA_LOG.Error("数据解密失败!", zap.Any("err", err))
        response.Fail(c)
        return
    }

    // 更新手机号
    userService.UpdatePhoneNumber(info.ID, mobile.PhoneNumber)

    var res = wxmpRes.WXPhoneNumberAuthResponse{
        Phone: mobile.PhoneNumber,
    }
    response.OkWithDetailed(&res, "登录成功", c)
}
