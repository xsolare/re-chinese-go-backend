package models

type Final struct {
	Id   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(8);not null"`
	Pos  uint8  `json:"pos"`
}

func (s *Final) TableName() string {
	return "finals"
}
