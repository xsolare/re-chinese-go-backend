package models

type HieroglyphKey struct {
	Id           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	HieroglyphId uint       `gorm:"not null" json:"hieroglyph_id"`
	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphKey) TableName() string {
	return "hieroglyphic_keys"
}
