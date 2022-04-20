package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type WordCollectionService struct{}

/// 																	///

func (wordCollectionService *WordCollectionService) CollectionById(id string) (err error, model []response.WordCollection) {
	var data []response.WordCollection
	db := global.GV_DB.Model(&models.WordCollection{})

	err = db.Raw(`
		select  w.id 				as id,
				w.name				as hieroglyphics,
				w.pinyin			as pinyin,
				w.part_of_speech_id as part_of_speech_id,
				wt.translate		as translate,
				w.hsk				as hsk
		from words_in_collections wic
		join words w on w.id = wic.word_id
		left join words_translates wt on wt.word_id = w.id
		where wic.word_collection_id = ?;
	`, id).Scan(&data).Error
	//? -- (cost=34.34..53.77 rows=9 width=114)

	return err, data
}
