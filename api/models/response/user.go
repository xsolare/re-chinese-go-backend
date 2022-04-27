package response

import (
	"time"

	"github.com/xsolare/re-chinese-go-backend/api/models"
)

type User struct {
	Id        uint          `json:"id"`
	Username  string        `json:"username"`
	Avatar    string        `json:"avatar"`
	UpdatedAt time.Time     `json:"updatedAt"`
	CreatedAt time.Time     `json:"createdAt"`
	LastLogin time.Time     `json:"lastLogin"`
	Roles     []models.Role `json:"roles" gorm:"many2many:users_roles;"`
}

type UserJwt struct {
	User User   `json:"user"`
	Jwt  string `json:"jwt"`
}
