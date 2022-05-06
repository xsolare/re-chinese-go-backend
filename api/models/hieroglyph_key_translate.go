package models

type HieroglyphKeyTranslate struct {
	Translate       string   `json:"translate" gorm:"type:varchar(60);unique;not null;primaryKey"`
	HieroglyphKeyId uint     `json:"hieroglyphKeyId" gorm:"primaryKey"`
	LanguageId      uint     `json:"languageId" gorm:"primaryKey"`
	Language        Language `json:"language" gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphKeyTranslate) TableName() string {
	return "hieroglyphic_keys_translates"
}
