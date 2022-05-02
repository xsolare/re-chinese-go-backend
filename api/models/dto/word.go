package dto

import "github.com/xsolare/re-chinese-go-backend/api/models"

type Word struct {
	Name   string `json:"name"`
	Pinyin string `json:"pinyin"`
	Hsk    uint8  `json:"hsk"`
}

func (s *Word) Generate(model *models.Word) {
	model.Name = s.Name
	model.Pinyin = s.Pinyin
	model.Hsk = s.Hsk
}
