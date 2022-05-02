package service

import (
	"strings"

	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
	"gorm.io/gorm"
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

func (wordService *WordService) AddWord(req *dto.Word) (err error, model models.Word) {

	global.GV_DB.Transaction(func(tx *gorm.DB) error {
		req.Generate(&model)
		err = tx.Model(&models.Word{}).Omit("id").Create(&model).Error

		var wordHieroglyphic []models.WordHieroglyph
		wordHieroglyphicStr := strings.Split(req.Name, "")
		for index, element := range wordHieroglyphicStr {
			err, h := new(HieroglyphService).ByName(element)

			if err != nil {
				return err
			}

			wordHieroglyphic = append(wordHieroglyphic, models.WordHieroglyph{
				WordId:       int(model.Id),
				HieroglyphId: int(h.Id),
				Index:        uint(index + 1),
			})
		}

		err = tx.Model(&models.WordHieroglyph{}).Create(&wordHieroglyphic).Error

		if err != nil {
			return err
		}
		return nil
	})

	return err, model
}
