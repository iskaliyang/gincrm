package api

import (
	"gincrm/models"
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type CustomerApi struct {
	customerService *service.CustomerService
}

func NewCustomerApi() *CustomerApi {
	customerApi := CustomerApi{
		customerService: &service.CustomerService{},
	}
	return &customerApi
}

func (c *CustomerApi) GetList(context *gin.Context) {
	var param models.CustomerQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}

	param.Creator = int64(uid)
	customerList, rows, code := c.customerService.GetList(&param)
	response.PageResult(code, customerList, rows, context)
}

func (c *CustomerApi) Create(context *gin.Context) {
	var param models.CustomerCreateParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	code := c.customerService.Create(&param)
	response.Result(code, nil, context)
}

func (c *CustomerApi) Export(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	file, code := c.customerService.Export(int64(uid))
	if len(file) >= 0 && code != 0 {
		response.Result(code, nil, context)
		return
	}
	context.File(file)
}

func (c *CustomerApi) Delete(context *gin.Context) {
	var param models.CustomerDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	code := c.customerService.Delete(&param)
	response.Result(code, nil, context)
}

func (c *CustomerApi) SendMail(context *gin.Context) {
	var param models.CustomerSendMailParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	param.Uid = int64(uid)
	errCode := c.customerService.SendMail(&param)
	response.Result(errCode, nil, context)
}

func (c *CustomerApi) Update(context *gin.Context) {
	var param models.CustomerUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	errCode := c.customerService.Update(&param)
	response.Result(errCode, nil, context)
}

func (c *CustomerApi) GetInfo(context *gin.Context) {
	var param models.CustomerQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	customerInfo, errCode := c.customerService.GetInfo(&param)
	response.Result(errCode, customerInfo, context)
}

func (c *CustomerApi) GetOption(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	customerOption, errCode := c.customerService.GetOption(int64(uid))
	response.Result(errCode, customerOption, context)
}
