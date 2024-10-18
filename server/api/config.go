package api

import (
	"gincrm/models"
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type MailConfigApi struct {
	mailConfigService *service.MailConfigService
}

func NewMailConfigApi() *MailConfigApi {
	return &MailConfigApi{
		mailConfigService: service.NewMailConfigService(),
	}
}

func (m *MailConfigApi) Save(context *gin.Context) {
	var param models.MailConfigSaveParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	param.Creator = int64(uid)
	errCode := m.mailConfigService.Save(&param)
	if errCode == 0 {
		mailConfig, errCode := m.mailConfigService.GetInfo(int64(uid))
		response.Result(errCode, mailConfig, context)
		return
	}
	response.Result(errCode, nil, context)
}

func (m *MailConfigApi) Info(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	mailConfig, errCode := m.mailConfigService.GetInfo(int64(uid))
	response.Result(errCode, mailConfig, context)
}

func (m *MailConfigApi) Check(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	errCode := m.mailConfigService.Check(int64(uid))
	response.Result(errCode, nil, context)
}

func (m *MailConfigApi) UpdateStatus(context *gin.Context) {
	var param models.MailConfigStatusParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	param.Creator = int64(uid)
	errCode := m.mailConfigService.UpdateStatus(&param)
	response.Result(errCode, nil, context)
}

func (m *MailConfigApi) Delete(context *gin.Context) {
	var param models.MailConfigDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	errCode := m.mailConfigService.Delete(&param)
	response.Result(errCode, nil, context)
}
