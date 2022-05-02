package controllers

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	res "github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type WordCollectionController struct{}

/// 																	///

func Shuffle(vals []res.WordCollection) {
  r := rand.New(rand.NewSource(time.Now().Unix()))
  for len(vals) > 0 {
    n := len(vals)
    randIndex := r.Intn(n)
    vals[n-1].Id, vals[randIndex].Id = vals[randIndex].Id, vals[n-1].Id
    vals = vals[:n-1]
  }
}


func (r *WordCollectionController) GetById(c *gin.Context) {
	// counts := c.Query("counts")
	id := c.Param("id")

	err, data := wordCollectionService.CollectionById(id)
	// Shuffle(data)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

///																											//
