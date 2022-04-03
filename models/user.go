package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                    uint                    `gorm:"primaryKey;autoIncrement"			json:"id"`
	Username              string                  `gorm:"type:varchar(32);unique;not null" 	json:"user_name"`
	Password              string                  `gorm:"type:varchar(64);not null" 			json:"password"`
	Avatar                string                  `											json:"avatar"`
	CreatedAt             time.Time               `gorm:"autoCreateTime"		 				json:"created_at"`
	UpdatedAt             time.Time               `gorm:"autoCreateTime"		 				json:"updated_at"`
	LastLogin             time.Time               `gorm:"autoCreateTime"		 				json:"last_login"`
	Roles                 []Role                  `gorm:"many2many:users_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HieroglyphCollections []HieroglyphCollections `gorm:"type:foreignKey:AuthorId;references:Id" 	json:"user_id"`
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}

func (s *User) TableName() string {
	return "users"
}
