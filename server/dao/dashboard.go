package dao

import (
	"gincrm/global"
	"gincrm/models"
)

type DashboardDao struct {
}

func NewDashboardDao() *DashboardDao {
	return &DashboardDao{}
}

func (d *DashboardDao) Count(uid int64, days int) models.DashboardSummary {
	var ds models.DashboardSummary
	global.DB.Raw("select count(*) from customers where creator = ?", uid).Scan(&ds.Customers)
	global.DB.Raw("select count(*) from contracts where creator = ?", uid).Scan(&ds.Contracts)
	global.DB.Raw("select sum(amount) from contracts where creator = ?", uid).Scan(&ds.ContractAmount)
	global.DB.Raw("select count(*) from products where creator = ?", uid).Scan(&ds.Products)

	s := "industry as name, count(*) as value"
	global.DB.Model(&models.Customer{}).Select(s).Where("creator = ?", uid).Group("industry").Find(&ds.CustomerIndustry)

	return ds
}

func (d *DashboardDao) AmountSummary(start, end, uid int64) float64 {
	var as float64
	con := "created > ? and created < ? and status = 1 and creator = ?"
	global.DB.Table(CONTRACT).Where(con, start, end, uid).Pluck("amount", &as)
	return as
}
