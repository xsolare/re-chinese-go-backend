package models

type HieroglyphCollection struct {
	Id                        uint         `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name                      string       `gorm:"type:varchar(50);unique;not null" 	json:"name"`
	Description               string       `gorm:"type:varchar(180)" 				json:"description"`
	Hsk                       uint8        `gorm:"type:smallint;default:1"			json:"hsk"`
	AuthorId                  uint         `json:"author_id"`
	CategorieId               uint         `json:"categorie_id"`
	HieroglyphicInCollections []Hieroglyph `gorm:"many2many:hierogplyphic_in_collections;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphCollection) TableName() string {
	return "hieroglyphic_collections"
}
