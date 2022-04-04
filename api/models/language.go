package models

type Language struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name      string `gorm:"type:varchar(32);unique;not null" 	json:"name"`
	ShortName string `gorm:"type:varchar(16);unique;not null" 	json:"short_name"`
}

func (s *Language) TableName() string {
	return "languages"
}
