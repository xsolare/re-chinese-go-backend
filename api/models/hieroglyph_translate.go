package models

type HieroglyphTranslate struct {
	Id           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Translate    string     `json:"translate" gorm:"type:text;not null"`
	Priority     uint8      `json:"priority"`
	HieroglyphId uint       `json:"hieroglyphId" gorm:"not null"`
	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LanguageId   uint       `json:"languageId" gorm:"not null"`
	Language     Language   `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Description  string     `gorm:"type:text" 							json:"description"`
}

func (s *HieroglyphTranslate) TableName() string {
	return "hieroglyphic_translates"
}
