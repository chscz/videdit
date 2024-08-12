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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.1"})
	}

	sort.Slice(videos, func(i, j int) bool {
		return videos[i].Order < videos[j].Order
	})

	if err := vh.editor.ValidateRequest(videos); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := checkExistDir(vh.videoCfg.OutputFilePath); err != nil {
		return c.String(http.StatusBadRequest, "파일이 올바르지 않습니다.")
	}
	newVideoID := shortid.MustGenerate()
	newVideoPath := fmt.Sprintf("%s/%s", vh.videoCfg.OutputFilePath, newVideoID)
	os.Mkdir(newVideoPath, 0777)

	trimVideoIDs, err := vh.editor.TrimVideo(newVideoID, videos)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.2"})
	}
	if len(videos) != len(trimVideoIDs) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.3"})
	}
	if len(trimVideoIDs) == 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "성공"})
	}

	if err := vh.editor.ConcatVideo(newVideoID, videoReq.Ext, trimVideoIDs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.4"})
	}

	b, err := json.Marshal(videoReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.5"})
	}

	r := &model.VideoCreate{
		ID:        newVideoID,
		CreatedAt: time.Now(),
		Request:   string(b),
		FilePath:  vh.videoCfg.OutputFilePath,
	}

	if err := vh.repo.CreateVideoRequest(c.Request().Context(), r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "잘못된 요청입니다.6"})
	}

	filename := fmt.Sprintf("%s.%s", newVideoID, videoReq.Ext)
	return c.JSON(http.StatusOK, map[string]string{
		"id":        newVideoID,
		"file_name": filename,
		"url":       fmt.Sprintf("http://localhost:3000/download/%s", filename),
	})
}
