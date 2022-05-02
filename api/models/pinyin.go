package models

type Pinyin struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	InitialId   uint      `json:"initalId"`
	Initial     Initial   `json:"initial" gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FinalToneId uint      `json:"finalToneId" gorm:"not null"`
	FinalTone   FinalTone `json:"finalTone" gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Pinyin) TableName() string {
	return "pinyin"
}
