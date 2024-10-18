package dao

import (
	"gincrm/global"
	"gincrm/models"
	"time"
)

type ContractDao struct {
}

func NewContractDao() *ContractDao {
	return &ContractDao{}
}

func (c *ContractDao) Create(param *models.ContractCreateParam) error {
	contract := models.Contract{
		Name:                param.Name,
		Amount:              param.Amount,
		BeginTime:           param.BeginTime,
		OverTime:            param.OverTime,
		Remarks:             param.Remarks,
		Cid:                 param.Cid,
		ContractProductList: param.ContractProductList,
		Status:              param.Status,
		Creator:             param.Creator,
		Created:             time.Now().Unix(),
	}
	return global.DB.Create(&contract).Error
}

func (c *ContractDao) Update(param *models.ContractUpdateParam) error {
	contract := models.Contract{
		Id:                  param.Id,
		Name:                param.Name,
		Amount:              param.Amount,
		BeginTime:           param.BeginTime,
		OverTime:            param.OverTime,
		Remarks:             param.Remarks,
		Cid:                 param.Cid,
		ContractProductList: param.ContractProductList,
		Status:              param.Status,
		Updated:             time.Now().Unix(),
	}
	db := global.DB.Model(&contract).Select("*").Omit("id", "creator", "created")
	return db.Updates(&contract).Error
}

func (c *ContractDao) Delete(param *models.ContractDeleteParam) error {
	return global.DB.Delete(&models.Contract{}, param.Ids).Error
}

func (c *ContractDao) GetList(param *models.ContractQueryParam) ([]*models.ContractList, int64, error) {
	contractList := make([]*models.ContractList, 0)
	field := "contracts.id, contracts.name, contracts.amount, contracts.begin_time, contracts.over_time, customers.name as cname, contracts.remarks, contracts.status, contracts.created, contracts.updated"
	where := "inner join customers on contracts.cid = customers.id and contracts.creator = ?"
	raw := "select count(*) from contracts where creator = ?"

	// 分页查询
	offset := (param.Page.PageNum - 1) * param.Page.PageSize
	db := global.DB.Offset(offset).Limit(param.Page.PageSize).Table(CONTRACT).Select(field)

	var rows int64
	if param.Id != NumberNull {
		db.Joins(where+" and contracts.id = ?", param.Creator, param.Id)
		global.DB.Raw(raw+" and contracts.id = ?", param.Creator, param.Creator).Scan(&rows)
	} else {
		if param.Status != NumberNull {
			db.Joins(where+" and contracts.status = ?", param.Creator, param.Status)
			global.DB.Raw(raw+" and contracts.status = ?", param.Creator, param.Status).Scan(&rows)
		} else {
			db.Joins(where, param.Creator)
			global.DB.Raw(raw, param.Creator).Scan(&rows)
		}
	}
	if err := db.Scan(&contractList).Error; err != nil {
		return nil, NumberNull, nil
	}
	return contractList, rows, nil
}

func (c *ContractDao) GetListByUid(uid int64) ([]*models.ContractList, error) {
	contracts := make([]*models.ContractList, 0)
	s := "contracts.id, contracts.name, contracts.amount, contracts.begin_time, contracts.over_time, customers.name as cname, contracts.remarks, contracts.status, contracts.created, contracts.updated"
	j := "left join customers on contracts.cid = customers.id and contracts.creator = ?"
	err := global.DB.Table(CONTRACT).Select(s).Joins(j, uid).Scan(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (c *ContractDao) GetInfo(param *models.ContractQueryParam) (*models.ContractInfo, error) {
	contract := models.Contract{
		Id: param.Id,
	}
	contractInfo := models.ContractInfo{}
	err := global.DB.Table(CONTRACT).Where(&contract).First(&contractInfo).Error
	if err != nil {
		return nil, err
	}
	return &contractInfo, nil
}

func (c *ContractDao) GetAddedPList(id int64) (*models.Contract, error) {
	var contract models.Contract
	err := global.DB.Table(CONTRACT).Select("product_list").First(&contract, id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
