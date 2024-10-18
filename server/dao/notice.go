package dao

import (
	"gincrm/global"
	"gincrm/models"
	"time"
)

const (
	Read   = 1
	UnRead = 2
)

type NoticeDao struct {
}

func NewNoticeDao() *NoticeDao {
	return &NoticeDao{}
}

func (n *NoticeDao) Create(params *models.NoticeCreateParam) error {
	notice := models.Notice{
		Content: params.Content,
		Status:  UnRead,
		Creator: params.Creator,
		Created: time.Now().Unix(),
	}

	return global.DB.Create(&notice).Error
}

func (n *NoticeDao) GetList(uid int64) ([]*models.NoticeList, error) {
	noticeList := make([]*models.NoticeList, 0)
	db := global.DB.Table(NOTICES).Where("creator = ?", uid)
	if err := db.Order("status desc").Order("created desc").Find(&noticeList).Error; err != nil {
		return nil, err
	}
	return noticeList, nil
}

func (n *NoticeDao) GetUnReadCountByUid(uid int64) (models.UnReadNotice, error) {
	var unRead models.UnReadNotice
	raw := "select count(*) from notices where status = 2 and creator = ?"
	if err := global.DB.Raw(raw, uid).Scan(&unRead.Count).Error; err != nil {
		return unRead, err
	}
	return unRead, nil
}
