package service

import (
	"errors"

	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/global"
	"gorm.io/gorm"
)

type UserService struct{}

/// 																	///

func (userService *UserService) Users() (err error, model []models.User) {
	var data []models.User

	db := global.GV_DB.Model(&models.User{})

	err = db.Find(&data).Error

	return err, data
}

func (userService *UserService) SignUp(req dto.SignUp) (err error, model models.User) {
	var user models.User

	if !errors.Is(global.GV_DB.Model(&models.User{}).
		Where("username = ?", req.Username).
		First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("Cannot record user"), model
	}

	req.Generate(&user)
	user.SetPassword(req.Password)

	err = global.GV_DB.Model(&models.User{}).Omit("id").Create(&user).Error

	return err, user
}

func (userService *UserService) SignIn(req dto.SignIn) (err error, model models.User) {
	var user models.User
	db := global.GV_DB.Model(&models.User{})

	err = db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return errors.New("Cannot found user"), model
	}

	if user.Verify(req.Password) {
		return err, user
	} else {
		global.GV_LOG.Error("Password incorrect", err)
		return errors.New("Password incorrect"), model
	}

}
