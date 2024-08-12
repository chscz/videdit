package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/chscz/videdit/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

var (
	errBadRequestCreate      = errors.New("동영상 생성 요청이 올바르지 않습니다")
	errBadVideoRequest       = errors.New("요청 동영상이 올바르지 않습니다")
	errCreateOutputDirFailed = errors.New("/output 디렉토리 생성 실패하였습니다")
	errDatabaseSaveFailed    = errors.New("요청내역 저장을 실패하였습니다")
)

type VideoRequest struct {
	Videos []*model.VideoTrim
	Ext    string `json:"ext"`
}

func (vh *VideoHandler) CreateVideo(c echo.Context) error {
	// request 바인딩
	var videoReq VideoRequest
	if err := c.Bind(&videoReq); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewErrorToMap(err))
	}
	videos := videoReq.Videos
	if len(videos) == 0 {
		return c.JSON(http.StatusBadRequest, model.NewErrorToMap(errBadRequestCreate))
	}

	// 편집 순서 정렬
	sort.Slice(videos, func(i, j int) bool {
		return videos[i].Order < videos[j].Order
	})

	// 요청값이 유효한지 검사
	idx, err := vh.editor.ValidateRequest(videos)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewVideoEditorError(idx, err))
	}

	// output 디렉토리 체크 및 생성
	if err := checkDir(vh.videoCfg.OutputFilePath); err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewDetailErrorToMap(errCreateOutputDirFailed, err))
	}

	// videoID 생성, output 디렉토리 내 임시 디렉토리 생성
	newVideoID := shortid.MustGenerate()
	newVideoPath := fmt.Sprintf("%s/%s", vh.videoCfg.OutputFilePath, newVideoID)
	os.Mkdir(newVideoPath, 0777)

	// 동영상 편집(자르기)
	trimVideoIDs, idx, err := vh.editor.TrimVideo(newVideoID, videos)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewVideoEditorError(idx, err))
	}
	if len(videos) != len(trimVideoIDs) {
		return c.JSON(http.StatusBadRequest, model.NewErrorToMap(errBadVideoRequest))
	}
	if len(trimVideoIDs) > 1 {
		// 동영상 편집(붙이기)
		if err := vh.editor.ConcatVideo(newVideoID, videoReq.Ext, trimVideoIDs); err != nil {
			return c.JSON(http.StatusBadRequest, model.NewErrorToMap(errBadVideoRequest))
		}
	}

	// json 직렬화
	b, err := json.Marshal(videoReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewErrorToMap(err))
	}

	// 동영상 편집 요청 내역 저장
	saveVideoCreate := &model.VideoCreate{
		ID:        newVideoID,
		CreatedAt: time.Now(),
		Request:   string(b),
		FilePath:  vh.videoCfg.OutputFilePath,
	}
	if err := vh.repo.CreateVideoRequest(c.Request().Context(), saveVideoCreate); err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewDetailErrorToMap(errDatabaseSaveFailed, err))
	}

	// client로 응답
	filename := fmt.Sprintf("%s.%s", newVideoID, videoReq.Ext)
	return c.JSON(http.StatusOK, map[string]string{
		"id":        newVideoID,
		"file_name": filename,
		"url":       fmt.Sprintf("http://localhost:3000/download/%s", filename),
	})
}
