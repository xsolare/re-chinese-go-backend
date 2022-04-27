package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, http.StatusText(http.StatusOK), c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, http.StatusText(http.StatusOK), c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, http.StatusText(http.StatusBadRequest), c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusBadRequest, data, message, c)
}
