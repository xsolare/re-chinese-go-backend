package models

type WordCollection struct {
	Id                 uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name               string `json:"name" gorm:"type:varchar(50);unique;not null"`
	Hsk                uint8  `json:"hsk" gorm:"type:smallint;default:1"`
	AuthorId           uint   `json:"authorId"`
	CategorieId        uint   `json:"categorieId"`
	WordsInCollections []Word `gorm:"many2many:words_in_collections;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *WordCollection) TableName() string {
	return "words_collections"
}
