package models

type PartOfSpeech struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name string `gorm:"type:varchar(32);unique;not null" 	json:"name"`
}

func (s *PartOfSpeech) TableName() string {
	return "part_of_speech"
}
