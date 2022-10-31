package word

import "time"

type Word struct {
	ID          uint      `gorm:"primarykey"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	CreateTime  time.Time `gorm:"create_time"`
	UpdatedTime time.Time `gorm:"updated_time"`
}
