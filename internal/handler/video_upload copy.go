// package handler

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/chscz/videdit/internal/model"
// 	"github.com/chscz/videdit/internal/util"
// 	"github.com/labstack/echo/v4"
// 	"github.com/teris-io/shortid"
// )

// var (
// 	errBadFileRequest         = errors.New("파일이 올바르지 않습니다")
// 	errInvalidFileExtension   = errors.New("파일 확장자가 올바르지 않습니다")
// 	errCreateUploadDirFailed  = errors.New("/upload 디렉토리 생성 실패하였습니다")
// 	errCreateUploadFileFailed = errors.New("업로드 파일 생성을 실패하였습니다")
// 	errUploadListSaveFailed   = errors.New("업로드 내역 저장을 실패하였습니다")
// )

// func (vh *VideoHandler) UploadVideo(c echo.Context) error {
// 	// 요청에서 업로드 파일 추출
// 	uploadFile, err := c.FormFile("upload_file")
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.NewErrorToMap(errBadFileRequest))
// 	}

// 	// 지원가능 확장자 여부 체크
// 	if valid := util.CheckFileExtension(uploadFile.Filename); !valid {
// 		return c.JSON(http.StatusBadRequest, util.NewErrorToMap(errInvalidFileExtension))
// 	}

// 	// 업로드 디렉토리 체크 및 업로드 파일 생성
// 	uf, err := uploadFile.Open()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, util.NewErrorToMap(errBadFileRequest))
// 	}
// 	defer uf.Close()

// 	if err := util.CheckDir(vh.videoCfg.UploadFilePath); err != nil {
// 		return c.JSON(http.StatusInternalServerError, util.NewDetailErrorToMap(errCreateUploadDirFailed, err))
// 	}

// 	filePath := fmt.Sprintf("%s/%s", vh.videoCfg.UploadFilePath, uploadFile.Filename)
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, util.NewDetailErrorToMap(errCreateUploadFileFailed, err))
// 	}
// 	defer file.Close()

// 	if _, err = io.Copy(file, uf); err != nil {
// 		return c.JSON(http.StatusInternalServerError, util.NewDetailErrorToMap(errCreateUploadFileFailed, err))
// 	}

// 	// 업로드 내역 저장
// 	videoID := shortid.MustGenerate()
// 	saveUploadFile := &model.VideoUpload{
// 		ID:        videoID,
// 		CreatedAt: time.Now(),
// 		FileName:  uploadFile.Filename,
// 		FilePath:  vh.videoCfg.UploadFilePath,
// 	}
// 	if err := vh.repo.CreateVideoUpload(c.Request().Context(), saveUploadFile); err != nil {
// 		return c.JSON(http.StatusInternalServerError, util.NewDetailErrorToMap(errUploadListSaveFailed, err))
// 	}

// 	// 응답
// 	res := map[string]string{
// 		"id":        videoID,
// 		"file_name": uploadFile.Filename,
// 	}
// 	return c.JSON(http.StatusOK, res)
// }
