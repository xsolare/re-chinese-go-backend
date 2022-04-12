package models

type Initial struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"		json:"id"`
	Name string `gorm:"type:varchar(8);unique;not null" json:"name"`
	Pos  uint8  `json:"pos"`
}

func (s *Initial) TableName() string {
	return "initials"
}
