package models

type HieroglyphTranslate struct {
	Id           uint       `gorm:"primaryKey;autoIncrement"			json:"id"`
	Translate    string     `gorm:"type:varchar(8);unique;not null" 	json:"name"`
	Priority     uint8      `json:"name"`
	Description  string     `gorm:"type:text" 							json:"description"`
	HieroglyphId uint       `gorm:"not null" 							json:"hieroglyph_id"`
	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LanguageId   uint       `gorm:"not null"							json:"language_id"`
	Language     Language   `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphTranslate) TableName() string {
	return "hieroglyph_translates"
}
