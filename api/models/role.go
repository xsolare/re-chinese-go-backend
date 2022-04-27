package models

type RoleType string

const (
	ADMIN   = 1
	PREMIUM = 2
	STAFF   = 3
)

type Role struct {
	Id   int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name RoleType `json:"name" gorm:"type:role_type;unique"`
}

func (s *Role) TableName() string {
	return "roles"
}
