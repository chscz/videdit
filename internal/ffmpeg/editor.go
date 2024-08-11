package ffmpeg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chscz/videdit/internal/config"
	"github.com/chscz/videdit/internal/model"
	"github.com/teris-io/shortid"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type VideoEditor struct {
	cfg config.Video
}

func NewVideoEditor(cfg config.Video) *VideoEditor {
	return &VideoEditor{cfg: cfg}
}

func (ve *VideoEditor) TrimVideo(newVideoID string, videos []*model.VideoTrim) ([]string, error) {
	var trimVideoIDs []string
	for _, video := range videos {
		oriVideoFile := fmt.Sprintf("%s/%s", ve.cfg.UploadFilePath, video.VideoFileName)
		ext := filepath.Ext(video.VideoFileName)
		trimVideoID := shortid.MustGenerate()
		newVideoFile := fmt.Sprintf("%s/%s/%s%s", ve.cfg.OutputFilePath, newVideoID, trimVideoID, ext)
		if err := ffmpeg_go.Input(oriVideoFile, ffmpeg_go.KwArgs{"ss": video.TrimStart}).
			Output(newVideoFile, ffmpeg_go.KwArgs{"t": video.TrimEnd - video.TrimStart}).OverWriteOutput().Run(); err != nil {
			return nil, err
		}
		trimVideoIDs = append(trimVideoIDs, trimVideoID+ext)
	}
	return trimVideoIDs, nil
}

func (ve *VideoEditor) ConcatVideo(newVideoID string, trimVideoIDs []string) error {
	output := fmt.Sprintf("%s/%s.avi", ve.cfg.OutputFilePath, newVideoID)
	concatFile := fmt.Sprintf("%s/%s/%s.txt", ve.cfg.OutputFilePath, newVideoID, newVideoID)
	if err := createConcatListFile(concatFile, trimVideoIDs); err != nil {
		return err
	}

	if err := ffmpeg_go.Input(concatFile, ffmpeg_go.KwArgs{"f": "concat", "safe": "0"}).
		Output(output, ffmpeg_go.KwArgs{"c": "copy"}).
		Run(); err != nil {
		return err
	}

	os.RemoveAll(fmt.Sprintf("%s/%s", ve.cfg.OutputFilePath, newVideoID))
	return nil
}

func createConcatListFile(filename string, videos []string) error {
	content := ""
	for _, video := range videos {
		content += fmt.Sprintf("file '%s'\n", video)
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("could not write concat file: %w", err)
	}

	return nil
}
