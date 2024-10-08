package model

import "time"

type VideoUpload struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	FileName  string    `gorm:"column:file_name" json:"file_name"`
	FilePath  string    `gorm:"column:file_path" json:"file_path"`
}

func (VideoUpload) TableName() string {
	return "video_upload"
}
