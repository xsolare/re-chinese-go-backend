package models

type Pinyin struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"		json:"id"`
	InitialId   uint      `json:"inital_id"`
	Initial     Initial   `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FinalToneId uint      `gorm:"not null"						json:"final_tone_id"`
	FinalTone   FinalTone `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Pinyin) TableName() string {
	return "pinyin"
}
