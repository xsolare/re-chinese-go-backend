package sdk

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type Api struct {
	Context *gin.Context
	Errors  error
}

//^ Response Error
func (e Api) Error(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

//^ Response OK
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}
