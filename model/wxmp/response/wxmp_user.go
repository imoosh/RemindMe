package response

// 请求session_key
type GetWxMiniProgramSessionKeyResponse struct {
    SessionKey string `json:"session_key"`
    OpenId     string `json:"openid"`
    Token      string `json:"token"`
}

// 微信小程序获取用户信息登陆
type WXMiniProgramOauthResponse struct {
    Token string `json:"token"`
}

// 用户信息
type UserDataResponse struct {
    Avatar   string `json:"avatar"`
    NickName string `json:"nickname"`
    Group    struct {
        Image string `json:"image"`
        Name  string `json:"name"`
    } `json:"group"`
    Verification string `json:"verification"`
    Mobile       string `json:"mobile"`
}

// 微信小程序手机号码
type WXPhoneNumberAuthResponse struct {
    Phone string `json:"phone"`
}
