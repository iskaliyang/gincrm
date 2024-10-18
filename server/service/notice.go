package service

import (
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
)

type NoticeService struct {
	noticeDao *dao.NoticeDao
}

func (n *NoticeService) GetList(uid int64) ([]*models.NoticeList, int) {
	noticeList, err := n.noticeDao.GetList(uid)
	if err != nil {
		return nil, response.CodeFailed
	}
	return noticeList, response.CodeSuccess
}

func (n *NoticeService) GetUnReadCount(uid int64) (models.UnReadNotice, int) {
	unReadCount, err := n.noticeDao.GetUnReadCountByUid(uid)
	if err != nil {
		return unReadCount, response.CodeFailed
	}
	return unReadCount, response.CodeSuccess
}
