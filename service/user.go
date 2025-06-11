package service

import (
	"errors"
	"github.com/mrluzy/blueball/entity/database"
	"github.com/mrluzy/blueball/entity/request"
	"github.com/mrluzy/blueball/global"
	"github.com/mrluzy/blueball/utils"
	"github.com/mrluzy/blueball/utils/snowflake"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u *UserService) Register(req *request.Register) (database.User, error) {
	// 1.判断用户是否存在
	if !errors.Is(global.DB.Where("username = ?", req.Username).First(&database.User{}).Error, gorm.ErrRecordNotFound) {
		return database.User{}, errors.New("this username is already registered, please check the information you filled in, or retrieve your password")
	}

	// 2.创建用户实例
	user := database.User{
		UserID:   snowflake.GenID(),
		Username: req.Username,
		Password: utils.BcryptHash(req.Password),
		Email:    req.Email,
	}

	// 3.保存到数据库
	if err := global.DB.Create(&user).Error; err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (u *UserService) Login(req *request.Login) (database.User, error) {
	var user database.User

	// 查询用户是否存在
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return database.User{}, errors.New("user not found or database error")
	}

	// 校验密码
	if !utils.BcryptCheck(req.Password, user.Password) {
		return database.User{}, errors.New("incorrect password")
	}

	return user, nil
}
