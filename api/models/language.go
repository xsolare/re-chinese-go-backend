package models

type Language struct {
	Id        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"type:varchar(32);unique;not null"`
	ShortName string `json:"shortName" gorm:"type:varchar(16);unique;not null"`
}

func (s *Language) TableName() string {
	return "languages"
}
