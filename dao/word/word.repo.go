package word

import (
	"word-book/initiable"
)

type FindData struct {
	ID           uint    `filed:"id" abc:"" condition:"="`
	Tile         string  `filed:"title" abc:"" condition:"="`
	StartTimeLt  *int    `filed:"start_time" abc:"nil" condition:"<"`
	StartTimeGt  *int    `filed:"start_time" abc:"nil" condition:"="`
	StartTime    *int    `filed:"sart_time" abc:"nil" condition:"="`
	ShowTime     *int    `filed:"show_time" abc:"nil" condition:"="`
	ComplexityNt *string `filed:"complexity" abc:"nil" condition:"!="`

	// 其他操作
	Page  int    `filed:"complexity" abc:"0" condition:"="`
	Limit int    `filed:"complexity" abc:"0" condition:"="`
	Order string `filed:"complexity" abc:"" condition:"="`
}

func Find(data FindData) (words []Word) {
	db := initiable.GetDefaultGorm()

	if data.ID != 0 {
		db = db.Where("id = ?", data.ID)
	}
	if data.Tile != "" {
		db = db.Where("tile = ?", data.Tile)
	}

	if data.StartTime != nil {
		db = db.Where("start_time = ?", data.StartTime)
	}
	if data.StartTimeLt != nil {
		db = db.Where("start_time < ?", data.StartTimeLt)
	}
	if data.StartTimeGt != nil {
		db = db.Where("start_time > ?", data.StartTimeGt)
	}
	if data.ShowTime != nil {
		db = db.Where("show_time = ?", data.ShowTime)
	}
	if data.ComplexityNt != nil {
		db = db.Where("complexity != ?", data.ComplexityNt)
	}
	if data.Limit != 0 {
		db = db.Limit(data.Limit)
	}
	if data.Page != 0 {
		db = db.Offset((data.Page - 1) * data.Limit)
	}
	if data.Order == "" {
		data.Order = "ID ASC"
	}
	db.Order(data.Order).Find(&words)
	return
}

func SetTodayStartTime(words []Word, today int) {
	var ids []int64
	for _, word := range words {
		ids = append(ids, int64(word.ID))
	}
	db := initiable.GetDefaultGorm()
	db.Model(&Word{}).Where("id in ?", ids).Updates(Word{StartTime: int64(today)})
	db.Commit()
}

func UdateShowTime(word Word) {
	db := initiable.GetDefaultGorm()
	showTime := word.ShowTime + 1
	db.Model(&word).Select("ShowTime").Updates(Word{ShowTime: showTime})
	db.Commit()
}
