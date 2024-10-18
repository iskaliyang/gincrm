package api

import (
	"gincrm/response"
	"gincrm/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DashboardApi struct {
	dashboardService *service.DashboardService
}

func NewDashboardApi() *DashboardApi {
	dashboardApi := DashboardApi{
		dashboardService: &service.DashboardService{},
	}
	return &dashboardApi
}

func (d *DashboardApi) Summary(context *gin.Context) {
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	days, _ := strconv.Atoi(context.Query("daysRange"))
	if days < 7 || days > 30 {
		response.Result(response.CodeParamInvalid, nil, context)
		return
	}
	sum := d.dashboardService.Summary(int64(uid), days)
	response.Result(response.CodeSuccess, sum, context)
}
