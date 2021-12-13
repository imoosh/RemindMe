package request

type Watermark struct {
    AppID     string `json:"appid"`
    TimeStamp int64  `json:"timestamp"`
}

type WXUserInfo struct {
    OpenID    string    `json:"openId,omitempty"`
    NickName  string    `json:"nickName"`
    AvatarUrl string    `json:"avatarUrl"`
    Gender    int       `json:"gender"`
    Country   string    `json:"country"`
    Province  string    `json:"province"`
    City      string    `json:"city"`
    UnionID   string    `json:"unionId,omitempty"`
    Language  string    `json:"language"`
    Watermark Watermark `json:"watermark,omitempty"`
}

type ResUserInfo struct {
    UserInfo      WXUserInfo `json:"userInfo"`
    RawData       string     `json:"rawData"`
    Signature     string     `json:"signature"`
    EncryptedData string     `json:"encryptedData"`
    IV            string     `json:"iv"`
}

type Login struct {
    Code     string      `json:"code"`
    UserInfo ResUserInfo `json:"userInfo"`
}


// 请求session_key
type GetWxMiniProgramSessionKeyRequest struct {
    AutoLogin bool   `json:"autoLogin"`
    Code      string `json:"code"`
}

// 微信小程序获取用户信息登陆
type WXMiniProgramOauthRequest struct {
    RawData       string `json:"rawData"`
    EncryptedData string `json:"encryptedData"`
    Event         string `json:"event"`
    IV            string `json:"iv"`
    SessionKey    string `json:"session_key"`
    Signature     string `json:"signature"`
}

// 用户信息
type UserDataRequest struct {
    Avatar   string `json:"avatar"`
    NickName string `json:"nickname"`
    Group struct {
        Image string `json:"image"`
        Name  string `json:"name"`
    } `json:"group"`
    Verification string `json:"verification"`
    Mobile       string `json:"mobile"`
}

// 微信小程序手机号码
type WXPhoneNumberAuthRequest struct {
    EncryptedData string `json:"encryptedData"`
    Event         string `json:"event"`
    IV            string `json:"iv"`
    SessionKey    string `json:"session_key"`
}