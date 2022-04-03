package models

type Hsk struct {
	Id                    uint                    `gorm:"primaryKey;autoIncrement"		json:"id"`
	Lvl                   uint                    `gorm:"unique;not null"		json:"lvl"`
	HieroglyphCollections []HieroglyphCollections `gorm:"type:foreignKey:HskId;references:Id" 	json:"hsk_id"`
}

func (s *Hsk) TableName() string {
	return "hsk"
}
