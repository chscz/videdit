package ffmpeg

import (
	"fmt"
	"path/filepath"

	"github.com/chscz/videdit/internal/model"
	"github.com/teris-io/shortid"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func (ve *VideoEditor) TrimVideo(newVideoID string, videos []*model.VideoTrim) ([]string, int, error) {
	var trimVideoIDs []string
	for i, video := range videos {
		oriVideoFile := fmt.Sprintf("%s/%s", ve.cfg.UploadFilePath, video.VideoFileName)
		ext := filepath.Ext(video.VideoFileName)
		trimVideoID := shortid.MustGenerate()
		newVideoFile := fmt.Sprintf("%s/%s/%s%s", ve.cfg.OutputFilePath, newVideoID, trimVideoID, ext)
		if err := ffmpeg_go.Input(oriVideoFile, ffmpeg_go.KwArgs{"ss": video.TrimStart}).
			Output(newVideoFile, ffmpeg_go.KwArgs{"t": video.TrimEnd - video.TrimStart}).
			OverWriteOutput().Run(); err != nil {
			return nil, i + 1, err
		}
		trimVideoIDs = append(trimVideoIDs, trimVideoID+ext)
	}
	return trimVideoIDs, 0, nil
}
