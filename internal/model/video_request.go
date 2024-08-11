package model

import "time"

type VideoRequest struct {
	ID        string    `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Request   string    `gorm:"column:request"`
}

func (VideoRequest) TableName() string {
	return "video_request"
}
