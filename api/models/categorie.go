package models

type Categorie struct {
	Id                   uint                   `gorm:"primaryKey;autoIncrement" 						json:"id"`
	Name                 string                 `gorm:"type:varchar(32);unique;not null" 				json:"name"`
	HieroglyphCollection []HieroglyphCollection `gorm:"foreignKey:CategorieId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Categorie) TableName() string {
	return "categories"
}
