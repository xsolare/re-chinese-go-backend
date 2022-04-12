package models

type Final struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"		json:"id"`
	Name string `gorm:"type:varchar(8);not null" json:"name"`
	Pos  uint8  `json:"pos"`
}

func (s *Final) TableName() string {
	return "finals"
}
