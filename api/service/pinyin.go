package service

import (
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type PinyinService struct{}

/// 																	///

func (pinyinService *PinyinService) PinyinsByInitalAndFinal(inital string, final string) (err error, model []response.Pinyin) {
	var data []response.Pinyin
	db := global.GV_DB.Model(&models.Pinyin{})

	//* with inital as (
	//* 	select i.id,
	//* 		i.name
	//* 	from initials i
	//* 	where i.name = ?
	//* ), final as (
	//* 	select f.id
	//* 	from finals f
	//* 	where f.name = ?
	//* )
	//* select  pn.id 												 as pinyin_id,
	//* 		concat((select name from inital),ft.name) 			 as pinyin,
	//* 		(select id from inital) 							 as initial_id,
	//* 		(select f.id from finals f where f.id = ft.final_id) as final_id,
	//* 		ft.tone												 as tone
	//* from pinyin pn
	//* join finals_tone ft on ft.final_id = (select id from final)
	//* where pn.final_tone_id = ft.id
	//* and pn.initial_id = (select id from inital);
	//? -- (cost=38.27..93.3 rows=2 width=56)

	//* select  pn.id 												 as pinyin_id,
	//* 		concat(i.name,ft.name) 								 as pinyin,
	//* 		i.id 												 as initial_id,
	//* 		(select f.id from finals f where f.id = ft.final_id) as final_id,
	//* 		ft.tone												 as tone
	//* from pinyin pn
	//* join finals_tone ft on ft.id = pn.final_tone_id
	//* join initials i on i.id = pn.initial_id
	//* where pn.initial_id = (select id from initials i where i.name = ?)
	//* and pn.final_tone_id in (select id from finals_tone ft
	//* where ft.final_id = (select id from finals f where f.name = ?));
	//? -- (cost=38.51..102.12 rows=2 width=56)

	err = db.Raw(`
		with inital as (
			select i.id,
				   i.name
			from initials i
			where i.id = ?
			--or 
			--where i.name = ?
		), final as (
			select f.id
			from finals f
			where f.id = ?
			--or 
			--where f.name = ?
		)
		select  pn.id 												 as pinyin_id,
				concat((select name from inital),ft.name) 			 as pinyin,
				(select id from inital) 							 as initial_id,
				(select f.id from finals f where f.id = ft.final_id) as final_id,
				ft.tone												 as tone
		from pinyin pn
		left join finals_tone ft on ft.final_id = (select id from final)
		where pn.final_tone_id = ft.id
		and pn.initial_id = (select id from inital);
	`, inital, final).Scan(&data).Error

	return err, data
}

func (pinyinService *PinyinService) ByTone(tone string) (err error, model []response.Pinyin) {
	var data []response.Pinyin
	db := global.GV_DB.Model(&models.Pinyin{})

	err = db.Raw(`
		select  p.id 											as id,
				concat((SELECT name
						FROM initials i
						WHERE i.id = p.initial_id ), ft.name)   as pinyin,
				p.initial_id 									as initial_id,
				ft.final_id 									as final_id,
				p.final_tone_id 								as final_tone_id,
				ft.tone											as tone,
				f.pos											as final_pos,
				i.pos											as initial_pos
		from pinyin p 
		left join finals_tone ft on ft.id = p.final_tone_id
		left join finals 	 f	on f.id  = ft.final_id
		left join initials 	 i	on i.id  = p.initial_id
		where ft.tone = ?;
	`, tone).Scan(&data).Error

	return err, data
}

func (pinyinService *PinyinService) Full() (err error, model []response.Pinyin) {
	var data []response.Pinyin
	db := global.GV_DB.Model(&models.Pinyin{})

	err = db.Raw(`
		select  p.id 											as id,
				concat((SELECT name
						FROM initials i
						WHERE i.id = p.initial_id ), ft.name)   as pinyin,
				p.initial_id 									as initial_id,
				ft.final_id 									as final_id,
				p.final_tone_id 								as final_tone_id,
				ft.tone											as tone,
				f.pos											as final_pos,
				i.pos											as initial_pos
		from pinyin p 
		left join finals_tone ft on ft.id = p.final_tone_id
		left join finals 	 f	on f.id  = ft.final_id
		left join initials 	 i	on i.id  = p.initial_id;
	`).Scan(&data).Error

	return err, data
}

func (pinyinService *PinyinService) Initials() (err error, model []models.Initial) {
	var data []models.Initial
	db := global.GV_DB.Model(&models.Initial{})

	err = db.Find(&data).Error

	return err, data
}

func (pinyinService *PinyinService) Finals() (err error, model []models.Final) {
	var data []models.Final
	db := global.GV_DB.Model(&models.Final{})

	err = db.Find(&data).Error

	return err, data
}
