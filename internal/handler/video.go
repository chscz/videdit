package handler

import (
	"context"

	"github.com/chscz/videdit/internal/config"
	"github.com/chscz/videdit/internal/model"
)

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
	ValidateRequest(videos []*model.VideoTrim) (int, error)
	TrimVideo(newVideoID string, videos []*model.VideoTrim) ([]string, int, error)
	ConcatVideo(newVideoID, extension string, trimVideoIDs []string) error
}

func NewVideoHandler(repo VideoRepository, editor VideoEditor, videoCfg config.Video) *VideoHandler {
	return &VideoHandler{
		repo:     repo,
		editor:   editor,
		videoCfg: videoCfg,
	}
}
