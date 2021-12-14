package response

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type Response struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
    Msg  string      `json:"msg"`
    Time int64       `json:"time"`
}

const (
    ERROR   = 0
    SUCCESS = 1
    TOKEN_EXPIRE = 401
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
    // 开始时间
    c.JSON(http.StatusOK, Response{
        code,
        data,
        msg,
        time.Now().Unix(),
    })
}

func Ok(c *gin.Context) {
    Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
    Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
    Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
    Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
    Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
    Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
    Result(ERROR, data, message, c)
}

func TokenExpired(c *gin.Context)  {
    Result(TOKEN_EXPIRE, map[string]interface{}{}, "授权已过期", c)
}