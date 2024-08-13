package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/chscz/videdit/internal/util"
	"github.com/labstack/echo/v4"
)

var errFileNotFound = errors.New("파일을 찾을 수 없습니다")

func (vh *VideoHandler) DownloadVideo(c echo.Context) error {
	filename := c.Param("filename")
	filePath := fmt.Sprintf("%s/%s", vh.videoCfg.OutputFilePath, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.JSON(http.StatusNotFound, util.NewErrorToMap(errFileNotFound))
	}
	return c.Attachment(filePath, filename)
}
