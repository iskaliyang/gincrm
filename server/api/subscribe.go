package api

import (
	"gincrm/common"
	"gincrm/models"
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SubscribeApi struct {
	subscribeService *service.SubscribeService
}

func NewSubscribeApi() *SubscribeApi {
	subscribeApi := SubscribeApi{
		subscribeService: service.NewSubscribeService(),
	}
	return &subscribeApi
}

func (s *SubscribeApi) GetInfo(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if int64(uid) <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	subscribeInfo, code := s.subscribeService.GetInfo(int64(uid))
	response.Result(code, subscribeInfo, context)
}

func (s *SubscribeApi) Pay(context *gin.Context) {
	var param models.SubscribePayParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if int64(uid) <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}

	param.Uid = int64(uid)
	payUrl, code := s.subscribeService.Pay(param)
	response.Result(code, payUrl, context)
}

// PayBack 支付成功后的回调
func (s *SubscribeApi) PayBack(context *gin.Context) {
	notifyReq := common.GetAlipay().VerifySign(context.Request)
	errCode := s.subscribeService.PayBack(notifyReq.GetString("out_trade_no"))

	//模拟支付宝回调
	//no := "20241015164048369439"
	//errCode := s.subscribeService.PayBack(no)

	context.String(http.StatusOK, "%s", "success")
	response.Result(errCode, nil, context)
}
