package dto

import "github.com/xsolare/re-chinese-go-backend/api/models"

type Hieroglyph struct {
	Hieroglyph     string `json:"hieroglyph"`
	Hsk            uint8  `json:"hsk"`
	PinyinId       uint   `json:"pinyin_id"`
}

func (s *Hieroglyph) Generate(model *models.Hieroglyph) {
	model.Hieroglyph = s.Hieroglyph
	model.PinyinId = s.PinyinId
	model.Hsk = s.Hsk
}
