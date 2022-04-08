package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type HieroglyphCollectionController struct{}

/// 																	///

func (r *HieroglyphCollectionController) GetById(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

///																											//

// //TODO Hieroglyph ------------------------------------------------------------------------------------------------//
// //? Get by hieroglyph
// //? Post hieroglyph
// //? Post translate hieroglyph
// //? Update hieroglyph

// //* Get by hieroglyph
// //? payload
// {
// 	hieroglyph: '好'
// }

// ///																											//
// //* Post hieroglyph
// //? payload
// {
// 	hieroglyph: '好'
// 	hsk: 3
// 	pinyin_id: 1
// 	part_of_speech_id: 4
// }

// //* Post hieroglyph_translate
// //? payload
// {
// 	hieroglyph_id: 2
// 	translate: '...'
// 	priority: 1
// 	!description: ''
// 	pinyin_id: 1
// 	?language_id: 1
// }

// ///																											//
// //* Update hieroglyph
// //? payload
// {
// 	hieroglyph_id: 1
// 	hieroglyph: '好你'
// 	pinyin: 'nǐhǎo'
// 	hsk: 3
// 	part_of_speech_id: 4
// }

// ///																											//
// //TODO Pinyin ----------------------------------------------------------------------------------------------//
// //? Get by pinyin(none tone)

// //* Get by pinyin(none tone)
// //? payload
// {
// 	inital: 'h'
// 	final: 'ao'
// }
// select concat(i.name,ft.name), * from pinyin pn
// join finals_tone ft on ft.id = pn.final_tone_id
// join initials i on i.id = pn.initial_id
// where pn.initial_id = (select id from initials i where i.name = 'h')
// and pn.final_tone_id in (select id from finals_tone ft
// 						 where ft.final_id = (select id from finals f where f.name = 'ao'));
// ///																											//

// //TODO Word ------------------------------------------------------------------------------------------------//
// //? Get by hieroglyph
// //? Post word
// //? Update word

// //* Get by hieroglyph
// //? payload
// {
// 	hieroglyph: '好你'
// }

// ///																											//
// //* Post word
// //? payload
// {
// 	hieroglyph: '好你'
// 	pinyin: 'nǐhǎo'
// 	hsk: 3
// 	part_of_speech_id: 4
// }

// //* Post word_translate
// //? payload
// {
// 	word_id: <posted word id>
// 	translate: '...'
// 	priority: 1
// 	!description: ''
// 	pinyin_id: 1
// 	?language_id: 1
// }

// //* Post word_hieroglyph
// //? payload
// {
// 	hieroglyph: '好你' // [hieroglyph](loop)
// }

// 1 - 好 = 4
// 2 - 你 = 6

// //* Post words_hieroglyphic
// //? 好
// {
// 	index: 1
// 	word_id: <posted word id>
// 	hieroglyph_id: 4
// }
// //? 你
// {
// 	index: 2
// 	word_id: <posted word id>
// 	hieroglyph_id: 6
// }
// ///																											//
// //* Update word
// //? payload
// {
// 	word_id: 1
// 	hieroglyph: '好你'
// 	pinyin: 'nǐhǎo'
// 	hsk: 3
// 	part_of_speech_id: 4
// }
// ///																											//
// //TODO Word collection -------------------------------------------------------------------------------------//
// //? Get by id
// //? Get by name
// //? Post collection
// //? Put word in collection

// ///																											//
// //* Get by id
// //? payload
// {
// 	collection_id: 1
// }
// ///																											//

// //* Get by name
// //? payload
// {
// 	name: 'nice'
// }
// ///																											//

// //* Post collection
// //? payload
// {
// 	name: 'col'
// 	!hsk: 3
// 	!author_id: <user id who posted>
// 	!categorie_id: 1
// }

// ///																											//
// //* Put word into collection
// //? payload []
// {
// 	word_collection_id: 1
//     word_id: <[...word ids]>
// }
// ///																											//
