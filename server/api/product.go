package api

import (
	"gincrm/models"
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductApi struct {
	productService *service.ProductService
}

func NewProductApi() *ProductApi {
	productApi := ProductApi{
		productService: &service.ProductService{},
	}
	return &productApi
}

func (p *ProductApi) GetList(context *gin.Context) {
	var param models.ProductQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	productList, rows, errCode := p.productService.GetList(&param)
	response.PageResult(errCode, productList, rows, context)
}

func (p *ProductApi) GetInfo(context *gin.Context) {
	var param models.ProductQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	productInfo, errCode := p.productService.GetInfo(&param)
	response.Result(errCode, productInfo, context)
}

func (p *ProductApi) Export(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	file, errCode := p.productService.Export(int64(uid))
	if len(file) >= 0 && errCode != 0 {
		response.Result(errCode, nil, context)
		return
	}
	context.File(file)
}

func (p *ProductApi) Create(context *gin.Context) {
	var param models.ProductCreateParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	errCode := p.productService.Create(&param)
	response.Result(errCode, nil, context)
}

func (p *ProductApi) Update(context *gin.Context) {
	var param models.ProductUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	errCode := p.productService.Update(&param)
	response.Result(errCode, nil, context)
}

func (p *ProductApi) Delete(context *gin.Context) {
	var param models.ProductDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	errCode := p.productService.Delete(&param)
	response.Result(errCode, nil, context)
}
