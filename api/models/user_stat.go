package models

type UserStat struct {
	UserId       uint `json:"userId" gorm:"primaryKey"`
	User         User `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CorrectCount uint `json:"correctCount" gorm:"default:0"`
	WrongCount   uint `json:"wrongCount" gorm:"default:0"`
	TotalPoints  uint `json:"totalPoints" gorm:"default:0"`
}

func (s *UserStat) TableName() string {
	return "users_stats"
}
