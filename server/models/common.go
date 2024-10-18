package models

type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type MailParam struct {
	Smtp       string
	Port       int
	AuthCode   string
	Sender     string
	Subject    string
	Content    string
	Attachment string
	Receiver   string
}
