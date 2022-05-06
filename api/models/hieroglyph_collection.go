package models

type HieroglyphCollection struct {
	Id                        uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                      string       `json:"name" gorm:"type:varchar(50);unique;not null"`
	Description               string       `json:"description" gorm:"type:varchar(180)"`
	Hsk                       uint8        `json:"hsk" gorm:"type:smallint;default:1"`
	AuthorId                  uint         `json:"authorId"`
	CategorieId               uint         `json:"categorieId"`
	HieroglyphicInCollections []Hieroglyph `gorm:"many2many:hierogplyphic_in_collections;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphCollection) TableName() string {
	return "hieroglyphic_collections"
}
