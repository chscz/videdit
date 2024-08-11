package model

type VideoUpload struct {
	ID       string `gorm:"column:id;primaryKey"`
	FileName string `gorm:"column:file_name"`
}

func (VideoUpload) TableName() string {
	return "video_upload"
}
