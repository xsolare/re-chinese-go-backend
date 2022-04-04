package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models"
)

var hieroglyph_collection = []models.HieroglyphCollection{
	{Id: 1, Name: "asd", HskId: 1, AuthorId: 1, CategorieId: 1},
	{Id: 2, Name: "fgh", HskId: 2, AuthorId: 2, CategorieId: 2},
}

type HieroglyphCollectionController struct{}

/// 																	///

func (r *HieroglyphCollectionController) GetHieroglyphCollection(c *gin.Context) {
	if hieroglyph_collection != nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
