package models

type ToneType string

const (
	FIRST  ToneType = "FIRST"
	SECOND ToneType = "SECOND"
	THIRD  ToneType = "THIRD"
	FOURTH ToneType = "FOURTH"
	NONE   ToneType = "NONE"
)

type FinalTone struct {
	Id      uint     `gorm:"primaryKey;autoIncrement"			json:"id"`
	Name    string   `gorm:"type:varchar(8);unique;not null" 	json:"name"`
	Tone    ToneType `sql:"type:tone_type" 						json:"tone"`
	FinalId uint     `json:"final_id"`
	Final   Final    `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (s *FinalTone) TableName() string {
	return "finals_tone"
}
