package model

import "time"

type VideoCreate struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Request   string    `gorm:"column:request" json:"request"`
	FilePath  string    `gorm:"column:file_path" json:"file_path"`
}

func (VideoCreate) TableName() string {
	return "video_create"
}
