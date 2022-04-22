package models

type RoleType string

const (
	ADMIN   = 1
	PREMIUM = 2
	STAFF   = 3
)

type Role struct {
	Id   int      `gorm:"primaryKey;autoIncrement" 	json:"id"`
	Name RoleType `gorm:"type:role_type;unique" 	json:"name"`
}

func (s *Role) TableName() string {
	return "roles"
}
