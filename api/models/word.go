package models

type Word struct {
	Id             uint         `json:"id" 		gorm:"primaryKey;autoIncrement"`
	Name           string       `json:"name" 	gorm:"type:varchar(50);not null"`
	Pinyin         string       `json:"pinyin" 	gorm:"type:varchar(50);not null"`
	Hsk            uint8        `json:"hsk" 	gorm:"type:smallint;default:1"`
	PartOfSpeechId uint         `json:"part_of_speech_id"`
	Hieroglyphic   []Hieroglyph `gorm:"many2many:words_hieroglyphic;"`
}

func (s *Word) TableName() string {
	return "words"
}