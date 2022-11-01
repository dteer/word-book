package word

import (
	"word-book/run"
)

type FindData struct {
	ID          uint
	Tile        string
	StartTimeLt int
	StartTime   int   `default:"-1"`
	ShowTime    int64 `default:"-1"` // -1代表不介入查询

	// 其他操作
	Page  int
	Limit int
	Order string
}

func Find(data FindData) (words []Word) {
	db := run.GetDefaultGorm()
	if data.ID != 0 {
		db = db.Where("id = ?", data.ID)
	}
	if data.Tile != "" {
		db = db.Where("tile = ?", data.Tile)
	}

	if data.StartTime != -1 {
		db = db.Where("start_time = ?", data.StartTime)
	}
	if data.StartTimeLt != 0 {
		db = db.Where("start_time < ?", data.StartTimeLt)
	}
	if data.ShowTime != -1 {
		db = db.Where("show_time = ?", data.ShowTime)
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
	db.Find(&words).Order(data.Order)
	return
}

func SetTodayStartTime(words []Word, today int) {
	var ids []int64
	for _, word := range words {
		ids = append(ids, int64(word.ID))
	}
	db := run.GetDefaultGorm()
	db.Model(&Word{}).Where("id in ?", ids).Updates(Word{StartTime: int64(today)})
	db.Commit()
}

func UdateShowTime(word Word) {
	db := run.GetDefaultGorm()
	showTime := word.ShowTime + 1
	db.Model(&word).Select("ShowTime").Updates(Word{ShowTime: showTime})
	db.Commit()
}
