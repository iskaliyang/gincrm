package dao

import (
	"fmt"
	"gincrm/global"
	"gincrm/models"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (u *UserDao) Create(params *models.UserCreateParams) error {
	user := models.User{
		Email:    params.Email,
		Password: params.Password,
		Status:   1,
		Created:  time.Now().Unix(),
	}
	return global.DB.Create(&user).Error
}

func (u *UserDao) IsExists(email string) bool {
	rows := global.DB.Where("email = ?", email).First(&models.User{}).RowsAffected
	return rows != NumberNull
}

func (u *UserDao) GetUser(email string) (*models.User, error) {
	var user models.User
	err := global.DB.Table(USER).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetCode(email string) string {
	key := fmt.Sprintf("user:%s:code", email)
	return global.REDIS.Get(ctx, key).Val()
}

func (u *UserDao) GetUid(email string) (int64, error) {
	user, err := u.GetUser(email)
	if err != nil {
		return NumberNull, err
	}
	return user.Id, nil
}

func (u *UserDao) SetCode(code int, email string) error {
	key := fmt.Sprintf("user:%s:code", email)
	return global.REDIS.Set(ctx, key, strconv.Itoa(code), 10*time.Minute).Err()
}

func (u *UserDao) GetInfo(uid int64) (*models.UserPersonInfo, error) {
	var user models.UserPersonInfo
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(USER).Where("id = ?", uid).First(&user).Error; err != nil {
			return err
		}
		//var subscribe models.Subscribe
		//if err := tx.Table(SUBSCRIBE).Select("version").Where("uid = ?", uid).First(&subscribe).Error; err != nil {
		//	return err
		//}
		//user.Version = subscribe.Version
		return nil
	})
	return &user, err
}
