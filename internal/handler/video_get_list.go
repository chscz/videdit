package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (vh *VideoHandler) GetVideoList(c echo.Context) error {
	uploadVideos, err := vh.repo.GetUploadVideoList(c.Request().Context())
	if err != nil {
		//
	}

	createVideos, err := vh.repo.GetCreateVideoList(c.Request().Context())
	if err != nil {
		//
	}

	resp := map[string]interface{}{
		"uploadVideos": uploadVideos,
		"reqVideos":    createVideos,
	}
	return c.JSON(http.StatusOK, resp)
}
