package api

import (
	"fmt"
	"gincrm/models"
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type UserApi struct {
	userService *service.UserService
}

func NewUserApi() *UserApi {
	userApi := UserApi{
		userService: service.NewUserService(),
	}
	return &userApi
}

func (u *UserApi) VerifyCode(ctx *gin.Context) {
	var params models.UserVerifyCodeParams
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("VerifyCode %s", err)
		response.Result(response.CodeParamInvalid, nil, ctx)
		return
	}

	code := u.userService.VerifyCode(params.Email)
	response.Result(code, nil, ctx)
}

func (u *UserApi) Register(ctx *gin.Context) {
	var params models.UserCreateParams
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.CodeParamInvalid, nil, ctx)
		log.Printf("[error]UserApi:Register:%s", err)
		return
	}

	code := u.userService.Register(&params)
	response.Result(code, nil, ctx)
}

func (u *UserApi) Login(ctx *gin.Context) {
	var params models.UserLoginParam
	if err := ctx.ShouldBind(&params); err != nil {
		response.Result(response.CodeParamInvalid, nil, ctx)
		return
	}

	userInfo, code := u.userService.Login(&params)
	if userInfo == nil {
		response.Result(code, nil, ctx)
		return
	}

	response.Result(code, userInfo, ctx)
}

func (u *UserApi) GetInfo(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	userInfo, errCode := u.userService.GetInfo(int64(uid))
	response.Result(errCode, userInfo, context)
}
