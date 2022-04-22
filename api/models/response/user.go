package response

import "github.com/xsolare/re-chinese-go-backend/api/models"

type UserJwt struct {
	User models.User
	Jwt  string
}
