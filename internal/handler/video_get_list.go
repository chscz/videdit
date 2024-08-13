package handler

import (
	"errors"
	"net/http"

	"github.com/chscz/videdit/internal/util"
	"github.com/labstack/echo/v4"
)

var (
	errUploadListNotFound = errors.New("동영상 업로드 내역을 찾을 수 없습니다")
	errCreateListNotFound = errors.New("동영상 생성 내역을 찾을 수 없습니다")
)

func (vh *VideoHandler) GetVideoList(c echo.Context) error {
	// 동영상 업로드 내역 조회
	uploadVideos, err := vh.repo.GetUploadVideoList(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewErrorToMap(errUploadListNotFound))
	}
	// 동영상 생성 내역 조회
	createVideos, err := vh.repo.GetCreateVideoList(c.Request().Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewErrorToMap(errCreateListNotFound))
	}

	resp := map[string]interface{}{
		"uploadVideos": uploadVideos,
		"reqVideos":    createVideos,
	}
	return c.JSON(http.StatusOK, resp)
}
