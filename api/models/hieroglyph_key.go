package models

type HieroglyphKey struct {
	Id           uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	HieroglyphId uint       `json:"hieroglyph_id" gorm:"not null"`
	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphKey) TableName() string {
	return "hieroglyphic_keys"
}
