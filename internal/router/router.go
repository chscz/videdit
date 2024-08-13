package router

import (
	"github.com/chscz/videdit/internal/handler"
	"github.com/labstack/echo/v4"
)

func InitRouter(vh *handler.VideoHandler) *echo.Echo {
	e := echo.New()

	e.Static("/", "public")

	e.POST("/upload", vh.UploadVideo)
	e.POST("/create_video", vh.CreateVideo)
	e.GET("/download/:filename", vh.DownloadVideo)
	e.GET("/get_video_list", vh.GetVideoList)

	vh.StartUploadJob()

	// e.POST("/test", test)

	return e
}
