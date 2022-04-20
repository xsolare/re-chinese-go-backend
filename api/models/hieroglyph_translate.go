package models

type HieroglyphTranslate struct {
	Id           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Translate    string     `gorm:"type:text;not null" json:"translate"`
	Priority     uint8      `json:"priority"`
	HieroglyphId uint       `gorm:"not null" json:"hieroglyph_id"`
	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LanguageId   uint       `gorm:"not null" json:"language_id"`
	Language     Language   `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Description  string     `gorm:"type:text" 							json:"description"`
}

func (s *HieroglyphTranslate) TableName() string {
	return "hieroglyphic_translates"
}
