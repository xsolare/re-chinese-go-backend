package models

type PartOfSpeech struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name       string `gorm:"type:varchar(32);unique;not null" 	json:"name"`
	Word       Word
	Hieroglyph Hieroglyph
}

func (s *PartOfSpeech) TableName() string {
	return "parts_of_speech"
}
