package ffmpeg

import (
	"fmt"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func (ve *VideoEditor) ConcatVideo(newVideoID, extension string, trimVideoIDs []string) error {
	output := fmt.Sprintf("%s/%s.%s", ve.cfg.OutputFilePath, newVideoID, extension)
	concatFile := fmt.Sprintf("%s/%s/%s.txt", ve.cfg.OutputFilePath, newVideoID, newVideoID)
	if err := createConcatListFile(concatFile, trimVideoIDs); err != nil {
		return err
	}

	if err := ffmpeg_go.Input(concatFile, ffmpeg_go.KwArgs{"f": "concat", "safe": "0"}).
		Output(output, ffmpeg_go.KwArgs{"c": "copy", "vsync": "vfr"}).
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
