package model

type VideoTrim struct {
	Order         int     `json:"order"`
	VideoID       string  `json:"videoId"`
	VideoFileName string  `json:"videoFileName"`
	TrimStart     float64 `json:"trimStart"`
	TrimEnd       float64 `json:"trimEnd"`
}
