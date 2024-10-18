package service

import (
	"encoding/json"
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
	"time"
)

type SubscribeService struct {
	subscribeDao *dao.SubscribeDao
	noticeDao    *dao.NoticeDao
}

func NewSubscribeService() *SubscribeService {
	subscribeService := SubscribeService{
		subscribeDao: dao.NewSubscribeDao(),
		noticeDao:    dao.NewNoticeDao(),
	}
	return &subscribeService
}

func (s *SubscribeService) GetInfo(uid int64) (*models.SubscribeInfo, int) {
	si, err := s.subscribeDao.GetInfo(uid)
	if err != nil {
		return nil, response.CodeFailed
	}

	// 判断用户订阅是否过期
	if si.Version == 2 && time.Now().Unix() > int64(si.Expired) {
		if err := s.subscribeDao.UpdateVersion(uid, 1); err != nil {
			return nil, response.CodeFailed
		}
	}
	subscribeInfo, err := s.subscribeDao.GetInfo(uid)
	if err != nil {
		return nil, response.CodeFailed
	}

	return subscribeInfo, response.CodeSuccess
}

func (s *SubscribeService) Pay(param models.SubscribePayParam) (*models.SubscribePayUrl, int) {
	// 构建订单支付信息
	tradeNo := common.GetAlipay().GenTradeNo()
	totalAmount := float64(param.Duration) * 0.6
	payUrl, err := common.GetAlipay().PagePay(tradeNo, totalAmount)
	if err != nil {
		return nil, response.CodeFailed
	}

	order := models.SubscribePayOrder{
		Uid:      param.Uid,
		TradeNo:  tradeNo,
		Duration: param.Duration,
	}

	if err := s.subscribeDao.SetOrder(tradeNo, &order); err != nil {
		return nil, response.CodeFailed
	}

	subscribePayUrl := models.SubscribePayUrl{
		PayUrl: payUrl,
	}
	return &subscribePayUrl, response.CodeSuccess
}

// PayBack 支付成功回调
func (s *SubscribeService) PayBack(outTradeNo string) int {

	// 获取订单信息
	var order models.SubscribePayOrder
	orderJson, err := s.subscribeDao.GetOrder(outTradeNo)
	if err != nil {
		return response.CodeFailed
	}
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		return response.CodeFailed
	}

	// 创建订阅信息
	duration := order.Duration * 24 * 60 * 60
	if !s.subscribeDao.IsExists(order.Uid) {
		subscribe := models.SubscribeCreateParam{
			Uid:     order.Uid,
			Version: 2,
			Expired: time.Now().Unix() + duration,
		}
		if err := s.subscribeDao.Create(&subscribe); err != nil {
			return response.CodeFailed
		}
	} else {
		si, err := s.subscribeDao.GetInfo(order.Uid)
		if err != nil {
			return response.CodeFailed
		}
		subscribe := models.SubscribeUpdateParam{
			Uid:     order.Uid,
			Version: 2,
			Expired: si.Expired + duration,
		}
		if err := s.subscribeDao.Update(&subscribe); err != nil {
			return response.CodeFailed
		}
	}

	// 订阅通知
	notice := models.NoticeCreateParam{
		Content: SubscribeProfessionalNotice,
		Creator: order.Uid,
	}
	_ = s.noticeDao.Create(&notice)

	return response.CodeSuccess
}
