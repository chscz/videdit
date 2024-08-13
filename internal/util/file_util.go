package util

import (
	"os"
	"path/filepath"
	"strings"
)

type fileFormat string

const (
	avi fileFormat = "avi"
	mov fileFormat = "mov"
	mp4 fileFormat = "mp4"
)

var validExtension = map[string]fileFormat{
	"avi": avi,
	"mov": mov,
	"mp4": mp4,
}

// 디렉토리 체크 후 없으면 생성
func CheckDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// 파일 확장자 체크
func CheckFileExtension(fileName string) bool {
	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	if _, exist := validExtension[ext]; exist {
		return true
	}
	return false
}
