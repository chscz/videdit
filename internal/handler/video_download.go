package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (vh *VideoHandler) DownloadVideo(c echo.Context) error {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("%s/%s", vh.videoCfg.OutputFilePath, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "파일을 찾을 수 없습니다.",
		})
	}
	return c.Attachment(filePath, filename)
}
