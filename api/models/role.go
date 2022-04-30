package models

type RoleType string

const (
	ADMIN_ID   = 1
	PREMIUM_ID = 2
	STAFF_ID   = 3
)

const (
	ADMIN   RoleType = "admin"
	PREMIUM RoleType = "premium"
	STAFF   RoleType = "staff"
)

type Role struct {
	Id   int      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name RoleType `json:"name" gorm:"unique" sql:"type:role_type"`
}

func (s *Role) TableName() string {
	return "roles"
}
