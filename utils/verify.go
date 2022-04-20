package utils

var (
	SignUpVerify = Rules{"UserName": {NotEmpty()}, "Password": {NotEmpty()}}
)
