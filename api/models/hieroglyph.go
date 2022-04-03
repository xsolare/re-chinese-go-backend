package models

type Hieroglyph struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"			json:"id"`
	Hieroglyph string `gorm:"type:varchar(8);unique;not null" 	json:"hieroglyph"`
	PinyinId   uint   `gorm:"not null"							json:"pinyin_id "`
	Pinyin     Pinyin `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Hieroglyph) TableName() string {
	return "hieroglyphic"
}
