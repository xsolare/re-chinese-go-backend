package models

type Hsk struct {
	Id                   uint                   `gorm:"primaryKey;autoIncrement"		json:"id"`
	Lvl                  uint                   `gorm:"unique;not null"					json:"lvl"`
	HieroglyphCollection []HieroglyphCollection `gorm:"foreignKey:HskId;references:Id"`
}

func (s *Hsk) TableName() string {
	return "hsk"
}
