package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type UserService struct{}

/// 																	///

func (userService *UserService) Users() (err error, model []models.User) {
	var data []models.User

	db := global.GV_DB.Model(&models.User{})

	err = db.Find(&data).Error

	return err, data
}
