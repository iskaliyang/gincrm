package initialize

import (
	"fmt"
	"gincrm/api"
	"gincrm/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	engine := gin.Default()
	route := engine.Group("/api")

	{
		// 用户模块
		newUserApi := api.NewUserApi()
		route.GET("/user/code", newUserApi.VerifyCode)
		route.POST("/user/register", newUserApi.Register)
		route.POST("/user/login", newUserApi.Login)
		route.GET("/user/info", newUserApi.GetInfo)

		// 通知模块
		newNoticeApi := api.NewNoticeApi()
		route.GET("/notice/list", newNoticeApi.GetList)
		route.GET("/notice/count", newNoticeApi.GetUnReadCount)

		// 订阅模块
		newSubscribeApi := api.NewSubscribeApi()
		route.GET("/subscribe/info", newSubscribeApi.GetInfo)
		route.POST("/subscribe/pay", newSubscribeApi.Pay)
		route.POST("/subscribe/payback", newSubscribeApi.PayBack)

		// 仪表盘模块
		newDashboardApi := api.NewDashboardApi()
		route.GET("/dashboard/summary", newDashboardApi.Summary)

		// 客户模块
		newCustomerApi := api.NewCustomerApi()
		route.GET("/customer/list", newCustomerApi.GetList)
		route.POST("/customer/create", newCustomerApi.Create)
		route.GET("/customer/export", newCustomerApi.Export)
		route.DELETE("/customer/delete", newCustomerApi.Delete)
		route.POST("/customer/send", newCustomerApi.SendMail)
		route.PUT("/customer/update", newCustomerApi.Update)
		route.GET("/customer/info", newCustomerApi.GetInfo)
		route.GET("/customer/option", newCustomerApi.GetOption)

		// 配置模块
		newMailConfigApi := api.NewMailConfigApi()
		route.PUT("/config/save", newMailConfigApi.Save)
		route.GET("/config/info", newMailConfigApi.Info)
		route.GET("/config/check", newMailConfigApi.Check)
		route.PUT("/config/status", newMailConfigApi.UpdateStatus)
		route.DELETE("/config/delete", newMailConfigApi.Delete)

		// 合同模块
		newContractApi := api.NewContractApi()
		route.POST("/contract/create", newContractApi.Create)
		route.GET("/contract/info", newContractApi.GetInfo)
		route.GET("/contract/list", newContractApi.GetList)
		route.GET("/contract/export", newContractApi.Export)
		route.POST("/contract/plist", newContractApi.GetProductList)
		route.PUT("/contract/update", newContractApi.Update)
		route.DELETE("/contract/delete", newContractApi.Delete)

		// 产品模块
		newProductApi := api.NewProductApi()
		route.GET("/product/list", newProductApi.GetList)
		route.GET("/product/info", newProductApi.GetInfo)
		route.GET("/product/export", newProductApi.Export)
		route.POST("/product/create", newProductApi.Create)
		route.PUT("/product/update", newProductApi.Update)
		route.DELETE("/product/delete", newProductApi.Delete)
	}

	addr := fmt.Sprintf(":%v", global.Config.Server.Port)
	_ = engine.Run(addr)
}
