package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	Status   int    `gorm:"status"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
}

func (t *User) TableName() string {
	return "user"
}

type UserVerifyCodeParams struct {
	Email string `form:"email" binding:"required,email"`
}

type UserCreateParams struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	Password string `json:"password" binding:"required"`
}

type UserLoginParam struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserInfo struct {
	Uid   int64  `json:"uid"`
	Token string `json:"token"`
}
type UserPersonInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Version int    `json:"version"`
}
