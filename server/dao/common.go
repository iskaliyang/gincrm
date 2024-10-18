package dao

import (
	"context"
	"gincrm/global"
	"gincrm/models"
)

const (
	USER       = "user"
	NOTICES    = "notices"
	SUBSCRIBE  = "subscribes"
	CONTRACT   = "contracts"
	CUSTOMER   = "customers"
	MailConfig = "mail_configs"
	PRODUCT    = "products"
	NumberNull = 0
	StringNull = ""
)

var ctx = context.Background()

// RestPage 分页查询
// page  设置起始页、每页条数,
// name  查询目标表的名称
// query 查询条件,
// dest  查询结果绑定的结构体,
// bind  绑定表结构对应的结构体
func restPage(page models.Page, name string, query interface{}, dest interface{}, bind interface{}) (int64, error) {
	if page.PageNum > 0 && page.PageSize > 0 {
		offset := (page.PageNum - 1) * page.PageSize
		global.DB.Offset(offset).Limit(page.PageSize).Table(name).Where(query).Find(dest)
	}
	res := global.DB.Table(name).Where(query).Find(bind)
	return res.RowsAffected, res.Error
}
