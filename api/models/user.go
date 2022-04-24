package models

import (
	"time"

	"github.com/xsolare/re-chinese-go-backend/utils"
	"gorm.io/gorm"
)

type User struct {
	Id                   uint                   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username             string                 `json:"username" gorm:"type:varchar(32);unique;not null"`
	Password             string                 `json:"password" gorm:"type:varchar(64);not null"`
	Salt                 string                 `json:"salt" gorm:"type:varchar(64);not null"`
	Avatar               string                 `json:"avatar"`
	UpdatedAt            time.Time              `json:"updatedAt" gorm:"autoCreateTime"`
	CreatedAt            time.Time              `json:"createdAt" gorm:"autoCreateTime"`
	LastLogin            time.Time              `json:"lastLogin" gorm:"autoCreateTime"`
	Roles                []Role                 `json:"roles" gorm:"many2many:users_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HieroglyphCollection []HieroglyphCollection `json:"hieroglyphCollection" gorm:"foreignKey:AuthorId;references:Id"`
}

func (m *User) AfterCreate(db *gorm.DB) error {
	userStat := UserStat{UserId: m.Id}
	db.Create(&userStat)
	return nil
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
	if db.Statement.Changed("LastLogin") {
		return nil
	}

	m.UpdatedAt = time.Now().Local()
	return nil
}

func (s *User) TableName() string {
	return "users"
}

//? SetPassword
func (u *User) SetPassword(value string) {
	u.Salt = utils.GenerateRandomKey16()
	u.Password = u.GetPasswordHash(value, u.Salt)
}

//? GetPasswordHash
func (u *User) GetPasswordHash(password string, salt string) string {
	passwordHash, err := utils.SetPassword(password, salt)
	if err != nil {
		return ""
	}
	return passwordHash
}

//? Verify
func (u *User) Verify(password string) bool {
	return u.GetPasswordHash(password, u.Salt) == u.Password
}
