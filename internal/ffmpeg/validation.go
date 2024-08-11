package ffmpeg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/chscz/videdit/internal/model"
)

type ProbeResult struct {
	Streams []struct {
		Duration string `json:"duration"`
	} `json:"streams"`
}

func (ve *VideoEditor) ValidateRequest(videos []*model.VideoTrim) error {
	for i, video := range videos {
		filePath := fmt.Sprintf("%s/%s", ve.cfg.UploadFilePath, video.VideoFileName)
		duration, err := getVideoDuration(filePath)
		if err != nil {
			return err
		}
		// 시작시간을 음수로 넣거나 값을 안 넣을 경우 0으로 대체
		if video.TrimStart <= -1 {
			video.TrimStart = 0
		}
		// 끝시간을 음수로 넣거나 값을 안 넣을 경우, 끝 시간이 원본 영상길이보다 클 경우 원본 영상 길이로 대체
		if video.TrimEnd <= -1 || video.TrimEnd > duration {
			video.TrimEnd = duration
		}
		if video.TrimStart >= video.TrimEnd || video.TrimStart >= duration {
			return &model.ValidateError{Message: fmt.Sprintf("%d 번째 영상 오류", i+1), Err: nil}
		}
	}
	return nil
}

func getVideoDuration(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "stream=duration", "-of", "json", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("failed to execute ffprobe command: %w", err)
	}

	var result ProbeResult
	if err := json.Unmarshal(out.Bytes(), &result); err != nil {
		return 0, fmt.Errorf("failed to parse ffprobe output: %w", err)
	}

	if len(result.Streams) == 0 {
		return 0, fmt.Errorf("no video streams found")
	}

	duration, err := strconv.ParseFloat(result.Streams[0].Duration, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration: %w", err)
	}

	return duration, nil
}
