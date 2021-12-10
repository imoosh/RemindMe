package wxmp

import (
    "RemindMe/global"
    "RemindMe/model/common/response"
    "RemindMe/model/wxmp/request"
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
func (b *UserApi) User(c *gin.Context) {
    // 获取
    token := c.Request.Header.Get("token")
    if token == "" {
        response.FailWithMessage("非法的token", c)
        return
    }

    var claims *request.CustomClaims
    if err, _ := jwtService.GetRedisJWT(token); err == redis.Nil {
        response.FailWithMessage("用户信息不存在", c)
        return
    } else if err != nil {
        response.Fail(c)
        return
    } else {
        j := utils.NewWxmpJWT()
        claims, err = j.ParseToken(token)
        if err != nil {
            response.Fail(c)
            return
        }
    }

    var res = wxmpRes.UserDataResponse{
        Avatar:   claims.Avatar,
        NickName: claims.Nickname,
    }

    response.OkWithData(&res, c)
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

    loginRes, err := weapp.Login(AppId, AppSecret, req.Code)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }

    var res = wxmpRes.GetWxMiniProgramSessionKeyResponse{
        SessionKey: loginRes.SessionKey,
    }

    response.OkWithData(&res, c)
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
        response.Fail(c)
        return
    }
    fmt.Println(info)

    token, err := b.tokenNext(c, info)
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

    var res = wxmpRes.UserDataResponse{
        Avatar:   "",
        NickName: "",
        Group: struct {
            Image string `json:"image"`
            Name  string `json:"name"`
        }{},
        Verification: "",
        Mobile:       "",
    }

    response.OkWithData(&res, c)
}

// 登录以后签发jwt
func (b *UserApi) tokenNext(c *gin.Context, user *weapp.UserInfo) (token string, err error) {
    j := &utils.WxmpJWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
    claims := wxmpReq.CustomClaims{
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
            NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
            ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
            Issuer:    "7gcat",                                               // 签名的发行者
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
