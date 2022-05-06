package models

type HieroglyphKey struct {
	Id                    uint                     `json:"id" gorm:"primaryKey;autoIncrement"`
	Index                 uint                     `json:"index" gorm:"type:smallint;default:1"`
	Priority              uint8                    `json:"priority" gorm:"type:smallint;default:1"`
	AlternativeHieroglyph string                   `json:"alternativeHieroglyph" gorm:"type:varchar(60)"`
	HieroglyphId          uint                     `json:"hieroglyphId" gorm:"not null"`
	Hieroglyph            Hieroglyph               `json:"hieroglyph" gorm:"foreignkey:Id;references:hieroglyph_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Translate             []HieroglyphKeyTranslate `json:"translate" gorm:"foreignKey:hieroglyph_key_id;references:id"`
}

func (s *HieroglyphKey) TableName() string {
	return "hieroglyphic_keys"
}
