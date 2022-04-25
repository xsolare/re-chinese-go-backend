package models

type Categorie struct {
	Id                   uint                   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name                 string                 `json:"name" gorm:"type:varchar(32);unique;not null"`
	HieroglyphCollection []HieroglyphCollection `gorm:"foreignKey:CategorieId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *Categorie) TableName() string {
	return "categories"
}
