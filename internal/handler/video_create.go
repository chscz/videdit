package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/chscz/videdit/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

type VideoRequest struct {
	Videos []*model.VideoTrim
	Ext    string `json:"ext"`
}

func (vh *VideoHandler) CreateVideo(c echo.Context) error {
	var videoReq VideoRequest
	if err := c.Bind(&videoReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	videos := videoReq.Videos
	if len(videos) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[i].Order < videos[j].Order
	})

	if err := vh.editor.ValidateRequest(videos); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	os.Mkdir(vh.videoCfg.OutputFilePath, 0777)
	newVideoID := shortid.MustGenerate()
	newVideoPath := fmt.Sprintf("%s/%s", vh.videoCfg.OutputFilePath, newVideoID)
	os.Mkdir(newVideoPath, 0777)

	trimVideoIDs, err := vh.editor.TrimVideo(newVideoID, videos)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}
	if len(videos) != len(trimVideoIDs) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}
	if len(trimVideoIDs) == 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "성공"})
	}

	if err := vh.editor.ConcatVideo(newVideoID, trimVideoIDs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}

	b, err := json.Marshal(videoReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}

	r := &model.VideoRequest{
		ID:        newVideoID,
		CreatedAt: time.Now(),
		Request:   string(b),
	}

	if err := vh.repo.CreateVideoRequest(c.Request().Context(), r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다."})
	}

	// 영상링크?다운로드?

	return c.JSON(http.StatusOK, map[string]string{"status": "성공"})
}

func (VideoRequest) String() string {
	return ""
}
