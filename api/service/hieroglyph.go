package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type HieroglyphService struct{}

/// 																	///

func (hieroglyphService *HieroglyphService) ById(id string) (err error, model models.Hieroglyph) {
	var data models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})

	err = db.Where("id = ?", id).First(&data).Error

	return err, data
}

func (hieroglyphService *HieroglyphService) ByName(name string) (err error, model models.Hieroglyph) {
	var data models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})
	// err = db.Preload("Pinyin.Initial").
	// 		 Preload("Pinyin.FinalTone").
	// 		 Preload(clause.Associations).
	// 		 Where("hieroglyph = ?", name).
	// 		 First(&data).Error
	
	err = db.Where("hieroglyph = ?", name).
			 First(&data).Error


	return err, data
}

func (hieroglyphService *HieroglyphService) ByPinyin(pinyin string) (err error, model []models.Hieroglyph) {
	var data []models.Hieroglyph

	db := global.GV_DB.Model(&models.Hieroglyph{})

	err = db.Where("pinyin_id = ?", pinyin).Find(&data).Error

	return err, data
}

func (hieroglyphService *HieroglyphService) AddHieroglyph(req *dto.Hieroglyph) (err error, model models.Hieroglyph) {
	var data models.Hieroglyph

	req.Generate(&data)

	err = global.GV_DB.Model(&models.Hieroglyph{}).Omit("id").Create(&data).Error

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
