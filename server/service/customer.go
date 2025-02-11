package service

import (
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
	"strconv"
	"time"
)

const (
	Open  = 1
	Close = 2
)

type CustomerService struct {
	customerDao   *dao.CustomerDao
	mailConfigDao *dao.MailConfigDao
}

func (c *CustomerService) GetList(param *models.CustomerQueryParam) ([]*models.CustomerList, int64, int) {
	customerList, rows, err := c.customerDao.GetList(param)
	if err != nil {
		return nil, NumberNull, response.CodeFailed
	}
	return customerList, rows, response.CodeSuccess
}

func (c *CustomerService) Create(param *models.CustomerCreateParam) int {
	if c.customerDao.IsExists(param.Name, param.Creator) {
		return response.CodeCustomerHasExist
	}
	if err := c.customerDao.Create(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *CustomerService) Export(uid int64) (string, int) {
	customers, err := c.customerDao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.CodeFailed
	}
	excelRows := make([]models.CustomerExcelRow, 0)
	var row models.CustomerExcelRow
	for _, c := range customers {
		row.Name = c.Name
		row.Source = c.Source
		row.Phone = c.Phone
		row.Email = c.Email
		row.Industry = c.Industry
		row.Level = c.Level
		row.Remarks = c.Remarks
		row.Region = c.Region
		row.Address = c.Address
		if c.Status == 1 {
			row.Status = "已签约"
		}
		if c.Status == 2 {
			row.Status = "未签约"
		}
		row.Created = time.Unix(c.Created, 0).Format("2006-01-02")
		if c.Updated != 0 {
			row.Updated = time.Unix(c.Updated, 0).Format("2006-01-02")
		}
		excelRows = append(excelRows, row)
	}
	sheet := "客户信息"
	columns := []string{"名称", "客户来源", "手机号", "邮箱", "客户行业", "客户级别", "备注", "地区", "详细地址", "成交状态", "创建时间", "更新时间"}
	fileName := "customer_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.CodeFailed
	}
	return file, response.CodeSuccess
}

func (c *CustomerService) Delete(param *models.CustomerDeleteParam) int {
	if err := c.customerDao.Delete(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *CustomerService) SendMail(param *models.CustomerSendMailParam) int {
	mc, err := c.mailConfigDao.GetInfo(param.Uid)
	if err != nil {
		return response.CodeMailSendFailed
	}
	if mc.Status == Close {
		return response.CodeMailSendUnEnable
	}
	mail := models.MailParam{
		Smtp:       mc.Stmp,
		Port:       mc.Port,
		AuthCode:   mc.AuthCode,
		Sender:     mc.Email,
		Subject:    param.Subject,
		Content:    param.Content,
		Attachment: param.Attachment,
		Receiver:   param.Receiver,
	}
	if err := common.SendMailToCustomer(mail); err != nil {
		return response.CodeMailSendFailed
	}
	return response.CodeSuccess
}

func (c *CustomerService) Update(param *models.CustomerUpdateParam) int {
	if err := c.customerDao.Update(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *CustomerService) GetInfo(param *models.CustomerQueryParam) (*models.CustomerInfo, int) {
	customerInfo, err := c.customerDao.GetInfo(param)
	if err != nil {
		return nil, response.CodeFailed
	}
	return customerInfo, response.CodeSuccess
}

func (c *CustomerService) GetOption(uid int64) ([]*models.CustomerOption, int) {
	option, err := c.customerDao.GetOption(uid)
	if err != nil {
		return nil, response.CodeFailed
	}
	return option, response.CodeSuccess
}
