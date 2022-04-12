package models

type UserStat struct {
	UserId       uint `gorm:"primaryKey" 		json:"user_id"`
	User         User `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CorrectCount uint `gorm:"default:0"			json:"correct_count"`
	WrongCount   uint `gorm:"default:0"			json:"wrong_count"`
	TotalPoints  uint `gorm:"default:0"			json:"total_points"`
}

func (s *UserStat) TableName() string {
	return "users_stats"
}
