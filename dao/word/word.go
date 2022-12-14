package word

import "time"

type Word struct {
	ID              uint      `gorm:"primarykey"`
	Title           string    `gorm:"title"`
	PhoneticStymbol string    `gorm:"phonetic_stymbol"`
	Description     string    `gorm:"description"`
	StartTime       int64     `gorm:"start_time"`
	ShowTime        int64     `gorm:"show_time"`
	Complexity      int       `gorm:"complexity"`
	CreateTime      time.Time `gorm:"create_time"`
	UpdatedTime     time.Time `gorm:"updated_time"`
}
