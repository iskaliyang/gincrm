package models

type Notice struct {
	Id      int64  `gorm:"primaryKey"`
	Content string `gorm:"content"`
	Status  int    `gorm:"status"`
	Creator int64  `gorm:"creator"`
	Created int64  `gorm:"creator"`
	Updated int64  `gorm:"updated"`
}

type NoticeCreateParams struct {
	Content string `json:"content"`
	Creator int64  `gorm:"creator"`
}

type NoticeList struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	Created int64  `json:"created"`
}

type UnReadNotice struct {
	Count int `json:"count"`
}

type NoticeCreateParam struct {
	Content string `json:"content"`
	Creator int64  `gorm:"creator"`
}
