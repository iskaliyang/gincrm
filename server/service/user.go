package service

import (
	"fmt"
	"gincrm/common"
	"gincrm/dao"
	"gincrm/models"
	"gincrm/response"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	userDao   *dao.UserDao
	noticeDao *dao.NoticeDao
}

func NewUserService() *UserService {
	userService := UserService{
		userDao: dao.NewUserDao(),
	}
	return &userService
}

func (u *UserService) VerifyCode(email string) int {
	code := common.RandIntSixDigit()
	if err := u.userDao.SetCode(code, email); err != nil {
		return response.CodeFailed
	}
	content := fmt.Sprintf("您的本次验证码为: %v，切勿向他人透露。", code)
	err := common.SendEmail(email, content)
	if err != nil {
		return response.CodeVerityCodeInvalid
	}
	return response.CodeSuccess
}

func (u *UserService) Register(params *models.UserCreateParams) int {
	if u.userDao.IsExists(params.Email) {
		return response.CodeUserHasExist
	}

	code := u.userDao.GetCode(params.Email)
	if code != params.Code {
		return response.CodeVerityCodeInvalid
	}

	password, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.CodeFailed
	}

	params.Password = string(password)
	if err := u.userDao.Create(params); err != nil {
		return response.CodeFailed
	}

	uid, err := u.userDao.GetUid(params.Email)
	if err != nil {
		return response.CodeFailed
	}

	_ = u.noticeDao.Create(&models.NoticeCreateParam{
		Content: RegisterNotice,
		Creator: uid,
	})

	return response.CodeSuccess
}

func (u *UserService) Login(params *models.UserLoginParam) (*models.UserInfo, int) {
	if !u.userDao.IsExists(params.Email) {
		return nil, response.CodeUserNotExist
	}

	user, err := u.userDao.GetUser(params.Email)
	if err != nil {
		return nil, response.CodeFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, response.CodeUserEmailOrPassword
	}

	token, err := common.GenToken(user.Id)
	if err != nil {
		log.Printf("[error]Login:GenerateToken:%s", err)
		return nil, response.CodeFailed
	}

	userInfo := models.UserInfo{
		Uid:   user.Id,
		Token: token,
	}

	_ = u.noticeDao.Create(&models.NoticeCreateParam{
		Content: LoginNotice,
		Creator: userInfo.Uid,
	})

	return &userInfo, response.CodeSuccess
}

func (u *UserService) GetInfo(uid int64) (*models.UserPersonInfo, int) {
	userInfo, err := u.userDao.GetInfo(uid)
	if err != nil {
		return nil, response.CodeFailed
	}
	return userInfo, response.CodeSuccess
}
