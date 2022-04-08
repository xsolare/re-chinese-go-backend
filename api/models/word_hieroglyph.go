package models

// type WordHieroglyphic struct {
// 	Id           uint       `gorm:"primaryKey;autoIncrement"			json:"id"`
// 	Index        uint       `gorm:"type:smallint;default:1"			json:"index"`
// 	WordId       uint       `gorm:"not null" 							json:"word_id"`
// 	Word         Word       `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	HieroglyphId uint       `gorm:"not null" 							json:"hieroglyph_id"`
// 	Hieroglyph   Hieroglyph `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// }

// func (s *WordHieroglyphic) TableName() string {
// 	return "words_hieroglyphic"
// }

type WordHieroglyph struct {
	WordId       int  `gorm:"primaryKey"`
	HieroglyphId int  `gorm:"primaryKey"`
	Index        uint `gorm:"type:smallint;default:1"			json:"index"`
}

func (s *WordHieroglyph) TableName() string {
	return "words_hieroglyphic"
}
