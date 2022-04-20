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
	db := global.GV_DB.Model(&models.User{})

	if !errors.Is(db.Where("username = ?", req.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("Cannot record user"), model
	}

	req.SignUp(&user)

	user.SetPassword(req.Password)

	return err, user
}
