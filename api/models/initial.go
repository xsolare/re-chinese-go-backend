package models

type Initial struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(8);unique;not null"`
	Pos  uint8  `json:"pos"`
}

func (s *Initial) TableName() string {
	return "initials"
}
