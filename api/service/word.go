package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type WordService struct{}

/// 																	///

func (wordService *WordService) WordsByHieroglyph(hieroglyph string) (err error, model []response.Word) {
	var data []response.Word
	db := global.GV_DB.Model(&models.Word{})

	db.SetupJoinTable(&models.Word{}, "Hieroglyphic", &models.WordHieroglyph{})

	err = db.Select("id", "name", "pinyin", "hsk").Where("name LIKE ?", "%"+hieroglyph+"%").Find(&data).Error

	return err, data
}
