package models

type WordTranslate struct {
	Id          uint     `gorm:"primaryKey;autoIncrement"			json:"id"`
	Translate   string   `gorm:"type:varchar(8);unique;not null" 	json:"translate"`
	Priority    uint8    `json:"priority"`
	Description string   `gorm:"type:text" 							json:"description"`
	WordId      uint     `gorm:"not null" 							json:"word_id"`
	Word        Word     `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LanguageId  uint     `gorm:"not null"							json:"language_id"`
	Language    Language `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *WordTranslate) TableName() string {
	return "words_translates"
}
