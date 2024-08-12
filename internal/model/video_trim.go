package model

type VideoTrim struct {
	Order         int     `json:"order"`
	VideoID       string  `json:"video_id"`
	VideoFileName string  `json:"video_file_name"`
	TrimStart     float64 `json:"trim_start"`
	TrimEnd       float64 `json:"trim_end"`
}
