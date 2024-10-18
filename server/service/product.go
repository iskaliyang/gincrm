package service

import (
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
	"strconv"
	"time"
)

type ProductService struct {
	productDao *dao.ProductDao
}

func NewProductService() *ProductService {
	return &ProductService{
		productDao: dao.NewProductDao(),
	}
}

func (p *ProductService) Create(param *models.ProductCreateParam) int {
	if p.productDao.IsExists(param.Name, param.Creator) {
		return response.CodeProductHasExist
	}
	if err := p.productDao.Create(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (p *ProductService) Update(param *models.ProductUpdateParam) int {
	if err := p.productDao.Update(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (p *ProductService) Delete(param *models.ProductDeleteParam) int {
	if err := p.productDao.Delete(param); err != nil {
		return response.CodeFailed
	}
	return response.CodeSuccess
}

func (p *ProductService) GetList(param *models.ProductQueryParam) ([]*models.ProductList, int64, int) {
	productList, rows, err := p.productDao.GetList(param)
	if err != nil {
		return nil, NumberNull, response.CodeFailed
	}
	return productList, rows, response.CodeSuccess
}

func (p *ProductService) GetInfo(param *models.ProductQueryParam) (*models.ProductInfo, int) {
	productInfo, err := p.productDao.GetInfo(param)
	if err != nil {
		return nil, response.CodeFailed
	}
	return productInfo, response.CodeSuccess
}

func (p *ProductService) Export(uid int64) (string, int) {
	products, err := p.productDao.GetListByUid(uid)
	if err != nil {
		return StringNull, response.CodeFileExportFailed
	}
	excelRows := make([]models.ProductExcelRow, 0)
	var row models.ProductExcelRow
	for _, p := range products {
		row.Name = p.Name
		if p.Status == 1 {
			row.Status = "已上架"
		}
		if p.Status == 2 {
			row.Status = "已下架"
		}
		if p.Type == 1 {
			row.Type = "默认"
		}
		row.Unit = p.Unit
		row.Code = p.Code
		row.Price = p.Price
		row.Description = p.Description
		row.Created = time.Unix(p.Created, 0).Format("2006-01-02")
		if p.Updated != 0 {
			row.Updated = time.Unix(p.Updated, 0).Format("2006-01-02")
		}
		excelRows = append(excelRows, row)
	}
	sheet := "产品信息"
	columns := []string{"名称", "是否上下架", "类型", "单位", "编码", "价格", "描述", "创建时间", "更新时间"}
	fileName := "product_" + strconv.FormatInt(uid, 10)
	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
	if err != nil {
		return StringNull, response.CodeFileExportFailed
	}
	return file, response.CodeSuccess
}
