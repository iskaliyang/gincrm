package dao

import (
	"encoding/json"
	"gincrm/global"
	"gincrm/models"
	"time"
)

type SubscribeDao struct {
}

func NewSubscribeDao() *SubscribeDao {
	return &SubscribeDao{}
}

func (s *SubscribeDao) GetInfo(uid int64) (*models.SubscribeInfo, error) {
	var si models.SubscribeInfo
	if err := global.DB.Table(SUBSCRIBE).First(&si, "uid = ?", uid).Error; err != nil {
		return nil, err
	}
	return &si, nil
}

func (s *SubscribeDao) UpdateVersion(uid int64, v int) error {
	return global.DB.Model(&models.Subscribe{}).Where("uid = ?", uid).Update("version", v).Error
}

func (s *SubscribeDao) SetOrder(tradeNo string, param *models.SubscribePayOrder) error {
	order, _ := json.Marshal(&param)
	return global.REDIS.Set(ctx, tradeNo, string(order), time.Minute*30).Err()
}

func (s *SubscribeDao) GetOrder(tradeNo string) (string, error) {
	orderJson, err := global.REDIS.Get(ctx, tradeNo).Result()
	if err != nil {
		return StringNull, err
	}
	return orderJson, nil
}

func (u *SubscribeDao) IsExists(uid int64) bool {
	var subscribe models.Subscribe
	db := global.DB.Table(SUBSCRIBE).Where("uid = ?", uid).First(&subscribe)
	return db.RowsAffected != NumberNull
}

func (s *SubscribeDao) Create(param *models.SubscribeCreateParam) error {
	subscribe := models.Subscribe{
		Uid:     param.Uid,
		Version: param.Version,
		Expired: param.Expired,
		Created: time.Now().Unix(),
	}
	return global.DB.Table(SUBSCRIBE).Create(&subscribe).Error
}

func (s *SubscribeDao) Update(param *models.SubscribeUpdateParam) error {
	subscribe := models.Subscribe{
		Version: param.Version,
		Expired: param.Expired,
		Updated: time.Now().Unix(),
	}
	return global.DB.Model(&models.Subscribe{}).Where("uid = ?", param.Uid).Updates(&subscribe).Error
}
