package utils

var (
	SignUpVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	SignInVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
