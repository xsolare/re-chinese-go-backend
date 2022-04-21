package models

import (
	"time"

	"github.com/xsolare/re-chinese-go-backend/utils"
	"gorm.io/gorm"
)

type User struct {
	Id                   uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username             string                 `gorm:"type:varchar(32);unique;not null" json:"userName"`
	Password             string                 `gorm:"type:varchar(64);not null" json:"password"`
	Salt                 string                 `gorm:"type:varchar(64);not null" json:"salt"`
	Avatar               string                 ``
	CreatedAt            time.Time              `gorm:"autoCreateTime"`
	UpdatedAt            time.Time              `gorm:"autoCreateTime"`
	LastLogin            time.Time              `gorm:"autoCreateTime"`
	Roles                []Role                 `gorm:"many2many:users_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HieroglyphCollection []HieroglyphCollection `gorm:"foreignKey:AuthorId;references:Id"`
}

func (m *User) AfterCreate(db *gorm.DB) error {
	userStat := UserStat{UserId: m.Id}
	db.Create(&userStat)
	return nil
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
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
