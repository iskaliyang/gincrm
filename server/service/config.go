package service

import (
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
)

type MailConfigService struct {
	mailConfigDao *dao.MailConfigDao
}

func NewMailConfigService() *MailConfigService {
	return &MailConfigService{
		mailConfigDao: dao.NewMailConfigDao(),
	}
}

func (m *MailConfigService) Save(param *models.MailConfigSaveParam) int {
	if err := m.mailConfigDao.Save(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (m *MailConfigService) GetInfo(uid int64) (*models.MailConfigInfo, int) {
	if !m.mailConfigDao.IsExists(uid) {
		return nil, response.CodeMailConfigUnSet
	}
	mc, err := m.mailConfigDao.GetInfo(uid)
	if err != nil {
		return nil, response.CodeFailed
	}
	mailConfig := models.MailConfigInfo{
		Id:       mc.Id,
		Stmp:     mc.Stmp,
		Port:     mc.Port,
		AuthCode: mc.AuthCode,
		Email:    mc.Email,
		Status:   mc.Status,
	}
	return &mailConfig, response.CodeSuccess
}

func (m *MailConfigService) Check(uid int64) int {
	mc, err := m.mailConfigDao.GetInfo(uid)
	if err != nil {
		return response.CodeMailConfigInvalid
	}
	if err = common.DialMail(mc.Stmp, mc.Port, mc.Email, mc.AuthCode); err != nil {
		return response.CodeMailConfigInvalid
	}
	return response.CodeSuccess
}

func (m *MailConfigService) UpdateStatus(param *models.MailConfigStatusParam) int {
	if !m.mailConfigDao.IsExists(param.Creator) {
		return response.CodeMailConfigUnSet
	}
	if err := m.mailConfigDao.UpdateStatus(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (m *MailConfigService) Delete(param *models.MailConfigDeleteParam) int {
	if err := m.mailConfigDao.Delete(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}
