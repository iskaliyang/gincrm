package initialize

import (
	"gincrm/global"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
)

func InitAlipay() {
	pay := global.Config.Alipay
	client, err := alipay.NewClient(pay.AppId, pay.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	client.SetReturnUrl(pay.ReturnURL).SetNotifyUrl(pay.NotifyURL)
	client.SetSignType("RSA2")
	global.Alipay = client
}
