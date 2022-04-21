package dto

import "github.com/xsolare/re-chinese-go-backend/api/models"

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *SignUp) Generate(model *models.User) {
	model.Username = s.Username
	model.Password = s.Password
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *SignIn) Generate(model *models.User) {
	model.Username = s.Username
	model.Password = s.Password
}
