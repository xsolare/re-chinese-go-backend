package models

type Hieroglyph struct {
	Id             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Hieroglyph     string `gorm:"type:varchar(8);not null" json:"hieroglyph"`
	Hsk            uint8  `gorm:"type:smallint;default:1" json:"hsk"`
	PartOfSpeechId uint   `gorm:"not null" json:"part_of_speech_id"`
	PinyinId       uint   `gorm:"not null" json:"pinyin_id"`
	Pinyin         Pinyin `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Hieroglyph) TableName() string {
	return "hieroglyphic"
}
