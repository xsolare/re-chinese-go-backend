package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/service"
	"github.com/xsolare/re-chinese-go-backend/global"
)

var operationRecordService = service.OperationRecordService{}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {

		var body []byte
		// var userId int //TODO

		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				global.GV_LOG.Error("read body from request error:", err)
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		//TODO
		// claims := jwtauth.ExtractClaims(c)
		// if claims["UserId"] != 0 {
		// 	userId = int(claims.ID)
		// } else {
		// 	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		// 	if err != nil {
		// 		userId = 0
		// 	}
		// 	userId = id
		// }
		record := models.OperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			// UserID: userId, //TODO
		}

		if strings.Index(c.GetHeader("Content-Type"), "multipart/form-data") > -1 {
			if len(record.Body) > 512 {
				record.Body = "File or Length out of limit"
			}
		}
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := operationRecordService.CreateOperationRecord(record);
		//
		err != nil {
			global.GV_LOG.Error("create operation record error:", err)
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
