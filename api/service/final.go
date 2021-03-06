package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type FinalService struct{}

/// 																	///

func (finalService *FinalService) Finals() (err error, model []models.Final) {
	var data []models.Final

	db := global.GV_DB.Model(&models.Final{})

	err = db.Find(&data).Error

	return err, data
}
