package request

import (
    "github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
    ID         uint   `json:"id"`
    OpenID     string `json:"openId"`
    Nickname   string `json:"nickName"`
    Gender     int    `json:"gender"`
    Province   string `json:"province"`
    Language   string `json:"language"`
    Country    string `json:"country"`
    City       string `json:"city"`
    Avatar     string `json:"avatarUrl"`
    UnionID    string `json:"unionId"`
    BufferTime int64

    jwt.StandardClaims
}
