package service

import (
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
	"strconv"

	"time"
)

type ContractService struct {
	contractDao *dao.ContractDao
	productDao  *dao.ProductDao
}

func NewContractService() *ContractService {
	return &ContractService{
		contractDao: dao.NewContractDao(),
		productDao:  dao.NewProductDao(),
	}
}

func (c *ContractService) Create(param *models.ContractCreateParam) int {
	if err := c.contractDao.Create(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *ContractService) Update(param *models.ContractUpdateParam) int {
	if err := c.contractDao.Update(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *ContractService) Delete(param *models.ContractDeleteParam) int {
	if err := c.contractDao.Delete(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (c *ContractService) GetList(param *models.ContractQueryParam) ([]*models.ContractList, int64, int) {
	contractList, rows, err := c.contractDao.GetList(param)
	if err != nil {
		return nil, NumberNull, response.CodeFailed
	}
	return contractList, rows, response.CodeSuccess
}

func (c *ContractService) GetInfo(param *models.ContractQueryParam) (*models.ContractInfo, int) {
	contractInfo, err := c.contractDao.GetInfo(param)
	if err != nil {
		return nil, response.CodeFailed
	}
	return contractInfo, response.CodeSuccess
}

func (c *ContractService) GetProductList(param *models.ContractQueryParam) ([]*models.Products, int) {
	if param.Id == 0 {
		products, err := c.productDao.GetListByIds(param.Pids)
		if err != nil {
			return nil, response.CodeFailed
		}
		return products, response.CodeSuccess
	}

	// 默认已添加的产品列表
	contract, err := c.contractDao.GetAddedPList(param.Id)
	if err != nil {
		return nil, response.CodeFailed
	}

	// 最终已添加的产品列表
	addedProductList := make([]*models.Products, 0)

	if len(param.Pids) == 0 {
		return addedProductList, response.CodeSuccess
	}

	addedPids := make([]int64, 0)
	for _, pid := range param.Pids {
		if len(*contract.ContractProductList) == 0 {
			addedPids = param.Pids
			break
		}
		for _, product := range *contract.ContractProductList {
			if pid == product.Id {
				addedProductList = append(addedProductList, product)
				continue
			}
			addedPids = append(addedPids, pid)
		}
	}
	products, err := c.productDao.GetListByIds(addedPids)
	if err != nil {
		return nil, response.CodeFailed
	}
	addedProductList = append(addedProductList, products...)
	return addedProductList, response.CodeSuccess
}

func (c *ContractService) Export(uid int64) (string, int) {
	contracts, err := c.contractDao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.CodeFailed
	}
	excelRows := make([]models.ContractExcelRow, 0)
	var row models.ContractExcelRow
	for _, c := range contracts {
		row.Name = c.Name
		row.Cname = c.Cname
		row.Amount = c.Amount
		row.BeginTime = c.BeginTime
		row.OverTime = c.OverTime
		row.Remarks = c.Remarks
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
	sheet := "合同信息"
	columns := []string{"合同名称", "客户名称", "合同金额", "合同开始时间", "合同结束时间", "备注", "签约状态", "创建时间", "更新时间"}
	fileName := "contract_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.CodeFailed
	}
	return file, response.CodeSuccess
}
