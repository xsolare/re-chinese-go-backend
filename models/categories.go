package models

type Categories struct {
	Id                    int                     `gorm:"primaryKey;autoIncrement" 	json:"id"`
	Name                  string                  `gorm:"type:varchar(32);unique;not null" 	json:"name"`
	HieroglyphCollections []HieroglyphCollections `gorm:"type:foreignKey:CategorieId;references:Id" 	json:"categories_id"`
}

// CREATE TABLE categories (
//     id                  serial                PRIMARY KEY,
//     name                VARCHAR ( 50 ) UNIQUE NOT NULL
// );
