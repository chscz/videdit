package handler

import (
	"context"
	"os"

	"github.com/chscz/videdit/internal/config"
	"github.com/chscz/videdit/internal/model"
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

type VideoHandler struct {
	repo     VideoRepository
	editor   VideoEditor
	videoCfg config.Video
}

type VideoRepository interface {
	CreateVideoUpload(ctx context.Context, file *model.VideoUpload) error
	CreateVideoRequest(ctx context.Context, req *model.VideoCreate) error
	GetUploadVideoList(ctx context.Context) ([]*model.VideoUpload, error)
	GetCreateVideoList(ctx context.Context) ([]*model.VideoCreate, error)
}

type VideoEditor interface {
	ValidateRequest(videos []*model.VideoTrim) error
	TrimVideo(newVideoID string, videos []*model.VideoTrim) ([]string, error)
	ConcatVideo(newVideoID, extension string, trimVideoIDs []string) error
}

func NewVideoHandler(repo VideoRepository, editor VideoEditor, videoCfg config.Video) *VideoHandler {
	return &VideoHandler{
		repo:     repo,
		editor:   editor,
		videoCfg: videoCfg,
	}
}

func checkExistDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}
