package models

type HieroglyphCollections struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"	json:"id"`
	Name     string `gorm:"type:varchar(50);unique;not null" 	json:"name"`
	HskId    uint   `json:"hsk_id"`
	AuthorId uint   `json:"author_id"`
	//	CategorieId uint   `json:"categorie_id"`
}
