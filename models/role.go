package models

type Role struct {
	Id   int    `gorm:"primaryKey;autoIncrement" 	json:"id"`
	Name string `gorm:"type:varchar(32);unique" 	json:"name"`
}

func (s *Role) TableName() string {
	return "roles"
}