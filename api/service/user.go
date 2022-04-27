package service

import (
	"errors"
	"time"

	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
	"gorm.io/gorm"
)

type UserService struct{}

/// 																	///

func (userService *UserService) Users() (err error, model []response.User) {
	var data []response.User

	db := global.GV_DB.Model(&models.User{})

	err = db.Preload("Roles").Find(&data).Error

	if err != nil {
		return errors.New("Cannot found user"), model
	}

	return err, data
}

func (userService *UserService) SignUp(req dto.SignUp) (err error, model response.User) {
	var user models.User

	if !errors.Is(global.GV_DB.Model(&models.User{}).
		Where("username = ?", req.Username).
		First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("Cannot record user"), model
	}

	req.Generate(&user)
	user.SetPassword(req.Password)

	err = global.GV_DB.Model(&models.User{}).Omit("id").Create(&user).Error

	if err != nil {
		return err, model
	}

	return userService.Auth(user.Id)
}

func (userService *UserService) SignIn(req dto.SignIn) (err error, model response.User) {
	var user models.User
	db := global.GV_DB.Model(&models.User{})

	err = db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return errors.New("Cannot found user"), model
	}

	if user.Verify(req.Password) {
		return userService.Auth(user.Id)
	} else {
		global.GV_LOG.Error("Password incorrect", err)
		return errors.New("Password incorrect"), model
	}
}

func (userService *UserService) RoleById(id uint) (err error, model []models.Role) {
	var user models.User
	db := global.GV_DB.Model(&models.User{})

	err = db.Preload("Roles").Where("id = ?", id).First(&user).Error

	if err != nil {
		return errors.New("Cannot found user"), model
	}

	return err, user.Roles
}

func (userService *UserService) UserById(id uint) (err error, model response.User) {
	var user response.User

	db := global.GV_DB.Model(&models.User{})

	err = db.Preload("Roles").Where("id = ?", id).First(&user).Error

	if err != nil {
		return errors.New("Cannot found user"), model
	}

	return err, user
}

func (userService *UserService) Auth(id uint) (err error, model response.User) {
	var user response.User
	db := global.GV_DB.Model(&models.User{})
	err = db.Preload("Roles").Where("id = ?", id).First(&user).Error

	if err != nil {
		return errors.New("Cannot found user"), model
	}

	global.GV_DB.Model(&models.User{}).Where("id = ?", user.Id).Update("LastLogin", time.Now().Local())

	return err, user
}
