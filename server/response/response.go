package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess             = 0
	CodeFailed              = 1
	CodeParamInvalid        = 2
	CodeUserHasExist        = 10001
	CodeUserNotExist        = 10002
	CodeUserEmailOrPassword = 10003
	CodeVerityCodeInvalid   = 10005
	CodeCustomerHasExist    = 20001
	CodeFileExportFailed    = 30001
	CodeProductHasExist     = 40001
	CodeMailConfigInvalid   = 50001
	CodeMailSendFailed      = 50002
	CodeMailSendUnEnable    = 50003
	CodeMailConfigUnSet     = 50004
)

var msg = map[int]string{
	CodeSuccess:             "success",
	CodeFailed:              "failed",
	CodeParamInvalid:        "param invalid",
	CodeUserHasExist:        "user has exist",
	CodeUserNotExist:        "user not exist",
	CodeUserEmailOrPassword: "user email or password error",
	CodeVerityCodeInvalid:   "verify code invalid",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Result(code int, data interface{}, ctx *gin.Context) {
	message := msg[code]
	ctx.JSON(http.StatusOK, Response{code, message, data})
}

type Page struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func PageResult(code int, data interface{}, rows int64, c *gin.Context) {
	message := msg[code]
	page := &Page{Total: rows, List: data}
	c.JSON(http.StatusOK, Response{code, message, page})
}
