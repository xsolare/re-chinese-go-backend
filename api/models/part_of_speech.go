package models

type PartOfSpeech struct {
	Id            uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string `json:"name" gorm:"type:varchar(32);unique;not null"`
	WordTranslate WordTranslate
}

func (s *PartOfSpeech) TableName() string {
	return "parts_of_speech"
}
