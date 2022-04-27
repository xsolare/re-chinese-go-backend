package models

type WordTranslate struct {
	Id          uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Translate   string   `json:"translate" gorm:"type:varchar(8);unique;not null"`
	Priority    uint8    `json:"priority"`
	Description string   `json:"description" gorm:"type:text"`
	WordId      uint     `json:"wordId" gorm:"not null"`
	Word        Word     `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LanguageId  uint     `json:"languageId" gorm:"not null"					`
	Language    Language `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *WordTranslate) TableName() string {
	return "words_translates"
}
