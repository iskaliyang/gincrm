package api

import (
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NoticeApi struct {
	noticeService *service.NoticeService
}

func NewNoticeApi() *NoticeApi {
	noticeApi := NoticeApi{
		noticeService: &service.NoticeService{},
	}
	return &noticeApi
}

func (n *NoticeApi) GetList(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	noticeList, errCode := n.noticeService.GetList(int64(uid))
	response.Result(errCode, noticeList, context)
}

func (n *NoticeApi) GetUnReadCount(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	if uid <= 0 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	unReadNotice, errCode := n.noticeService.GetUnReadCount(int64(uid))
	response.Result(errCode, unReadNotice, context)
}
