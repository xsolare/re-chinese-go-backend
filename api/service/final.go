package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type FinalService struct{}

/// 																	///

func (finalService *FinalService) Finals() (err error, model models.Final) {
	var data models.Final

	err = global.GV_DB.First(&data).Error

	if err != nil {
		return err, data
	}

	return err, data
}
