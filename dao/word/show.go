package word

import "time"

type Show struct {
	ID          uint64    `gorm:"primarykey"`
	WordID      uint64    `gorm:"word_id"`
	ShowTime    int64     `gorm:"show_time"`
	CreateTime  time.Time `gorm:"create_time"`
	UpdatedTime time.Time `gorm:"updated_time"`
}

func CreateUpdate()
