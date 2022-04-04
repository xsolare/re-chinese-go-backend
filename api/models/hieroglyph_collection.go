package models

type HieroglyphCollection struct {
	Id                     uint         `gorm:"primaryKey;autoIncrement"				json:"id"`
	Name                   string       `gorm:"type:varchar(50);unique;not null" 	json:"name"`
	HskId                  uint         `json:"hsk_id"`
	AuthorId               uint         `json:"author_id"`
	CategorieId            uint         `json:"categorie_id"`
	HieroglyphInCollection []Hieroglyph `gorm:"many2many:hierogplyph_in_collection;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *HieroglyphCollection) TableName() string {
	return "hieroglyph_collections"
}
