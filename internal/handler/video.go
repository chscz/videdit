package handler

import (
	"context"

	"github.com/chscz/videdit/internal/config"
	"github.com/chscz/videdit/internal/model"
)

const UploadFilePath = "./upload"

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
	CreateVideoRequest(ctx context.Context, req *model.VideoRequest) error
}

type VideoEditor interface {
	ValidateRequest(videos []*model.VideoTrim) error
	TrimVideo(newVideoID string, videos []*model.VideoTrim) ([]string, error)
	ConcatVideo(newVideoID string, trimVideoIDs []string) error
}

func NewVideoHandler(repo VideoRepository, editor VideoEditor, videoCfg config.Video) *VideoHandler {
	return &VideoHandler{
		repo:     repo,
		editor:   editor,
		videoCfg: videoCfg,
	}
}
