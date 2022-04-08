package models

type WordCollection struct {
	Id                 uint   `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name               string `gorm:"type:varchar(50);unique;not null" 	json:"name"`
	Hsk                uint8  `gorm:"type:smallint;default:1"			json:"hsk"`
	AuthorId           uint   `json:"author_id"`
	CategorieId        uint   `json:"categorie_id"`
	WordsInCollections []Word `gorm:"many2many:words_in_collections;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *WordCollection) TableName() string {
	return "words_collections"
}
