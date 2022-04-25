package models

type WordHieroglyph struct {
	WordId       int  `json:"wordId" gorm:"primaryKey"`
	HieroglyphId int  `json:"hieroglyphId" gorm:"primaryKey"`
	Index        uint `json:"index" gorm:"type:smallint;default:1"`
}

func (s *WordHieroglyph) TableName() string {
	return "words_hieroglyphic"
}
