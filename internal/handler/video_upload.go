package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chscz/videdit/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

func (vh *VideoHandler) UploadVideo(c echo.Context) error {
	uploadFile, err := c.FormFile("upload_file")
	if err != nil {
		return c.String(http.StatusBadRequest, "파일이 올바르지 않습니다.")
	}
	if valid := checkFileExtension(uploadFile.Filename); !valid {
		fmt.Println("불합격")
	}

	uf, err := uploadFile.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer uf.Close()

	if err := checkExistDir(vh.videoCfg.UploadFilePath); err != nil {
		return c.String(http.StatusBadRequest, "파일이 올바르지 않습니다.")
	}

	filePath := fmt.Sprintf("%s/%s", vh.videoCfg.UploadFilePath, uploadFile.Filename)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err = io.Copy(file, uf); err != nil {
		fmt.Println(err)
	}

	videoID := shortid.MustGenerate()
	saveUploadFile := &model.VideoUpload{
		ID:        videoID,
		CreatedAt: time.Now(),
		FileName:  uploadFile.Filename,
		FilePath:  vh.videoCfg.UploadFilePath,
	}
	if err := vh.repo.CreateVideoUpload(c.Request().Context(), saveUploadFile); err != nil {
		fmt.Println(err)
	}

	res := map[string]string{
		"id":        videoID,
		"file_name": uploadFile.Filename,
	}
	return c.JSON(http.StatusOK, res)
}

func checkFileExtension(fileName string) bool {
	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	if _, exist := validExtension[ext]; exist {
		return true
	}
	return false
}
