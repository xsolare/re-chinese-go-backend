package models

type Hieroglyph struct {
	Id             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Hieroglyph     string `json:"hieroglyph" gorm:"type:varchar(8);not null"`
	Hsk            uint8  `json:"hsk" gorm:"type:smallint;default:1"`
	PartOfSpeechId uint   `json:"partOfSpeechId" gorm:"not null"`
	PinyinId       uint   `json:"pinyinId" gorm:"not null"`
	Pinyin         Pinyin `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Hieroglyph) TableName() string {
	return "hieroglyphic"
}
