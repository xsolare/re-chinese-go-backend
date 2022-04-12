package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type HieroglyphService struct{}

/// 																	///

func (hieroglyphService *HieroglyphService) ByName(name string) (err error, model []models.Hieroglyph) {
	var data []models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})

	err = db.Where("name = ?", name).Find(&data).Error

	return err, data
}

func (hieroglyphService *HieroglyphService) AddHieroglyph(c *dto.Hieroglyph) (err error, model models.Hieroglyph) {
	var data models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})

	c.Generate(&data)

	err = db.Omit("id").Create(&data).Error

	global.GV_LOG.Warn("ID - ", data.Id)

	return err, data
}

func (hieroglyphService *HieroglyphService) DeleteHieroglyph(id string) (err error, model models.Hieroglyph) {
	var data models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})
	err = db.Where("id = ?", id).Delete(&data).Error

	global.GV_LOG.Warn("ID - ", data.Id)

	return err, data
}
