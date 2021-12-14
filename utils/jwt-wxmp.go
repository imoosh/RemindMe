package utils

import (
	"RemindMe/global"
	"RemindMe/model/wxmp/request"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type WxmpJWT struct {
	SigningKey []byte
}

var (
	WxmpTokenExpired     = errors.New("Token is expired")
	WxmpTokenNotValidYet = errors.New("Token not active yet")
	WxmpTokenMalformed   = errors.New("That's not even a token")
	WxmpTokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewWxmpJWT() *WxmpJWT {
	return &WxmpJWT{
		[]byte(global.Config.JWT.SigningKey),
	}
}

// 创建一个token
func (j *WxmpJWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *WxmpJWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.ConcurrencyControl.Do("WxmpJWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *WxmpJWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, WxmpTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, WxmpTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, WxmpTokenNotValidYet
			} else {
				return nil, WxmpTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, WxmpTokenInvalid

	} else {
		return nil, WxmpTokenInvalid

	}

}
