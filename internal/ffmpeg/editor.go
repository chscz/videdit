package ffmpeg

import (
	"github.com/chscz/videdit/internal/config"
)

type VideoEditor struct {
	cfg config.Video
}

func NewVideoEditor(cfg config.Video) *VideoEditor {
	return &VideoEditor{cfg: cfg}
}
