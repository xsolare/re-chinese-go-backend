package models

//? type ToneType string
//? const (
//? 	FIRST  ToneType = "FIRST"
//? 	SECOND ToneType = "SECOND"
//? 	THIRD  ToneType = "THIRD"
//? 	FOURTH ToneType = "FOURTH"
//? 	NONE   ToneType = "NONE"
//? )

type FinalTone struct {
	Id      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"type:varchar(8);not null"`
	Tone    uint   `json:"tone" gorm:"type:smallint;default:1"`
	FinalId uint   `json:"finalId"`
	Final   Final  `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	//? Tone    ToneType `sql:"type:tone_type" 				json:"tone"`
}

func (s *FinalTone) TableName() string {
	return "finals_tone"
}
