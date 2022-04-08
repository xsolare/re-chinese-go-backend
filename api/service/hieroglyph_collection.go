package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
)

type HieroglyphCollectionService struct{}

/// 																	///

func (hieroglyphCollectionService *HieroglyphCollectionService) HieroglyphCollections() (err error, model []models.HieroglyphCollection) {
	var hieroglyph_collection = []models.HieroglyphCollection{
		{Id: 1, Name: "asd", Hsk: 1, AuthorId: 1, CategorieId: 1},
		{Id: 2, Name: "fgh", Hsk: 2, AuthorId: 2, CategorieId: 2},
	}

	return nil, hieroglyph_collection
}
