package router

import (
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/chscz/videdit/internal/handler"
	"github.com/labstack/echo/v4"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func InitRouter(vh *handler.VideoHandler) *echo.Echo {
	e := echo.New()

	e.Static("/", "public")

	e.POST("/upload", vh.UploadVideo)
	e.POST("/create_video", vh.CreateVideo)
	e.POST("/test", test)

	e.GET("/download/:filename", vh.DownloadVideo)
	e.GET("/get_video_list", vh.GetVideoList)

	return e
}

// func test(c echo.Context) error {
// 	// d, err := ffmpeg.GetVideoDuration("./upload/green.mp4")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// fmt.Println(d)
// 	return nil
// }

func test(c echo.Context) error {
	videoA := "./test/red.mp4"
	videoB := "./test/pink.mp4"
	videos := []string{videoA, videoB}
	output := "./test/output.mp4"

	concatFile := "concat_list.txt"
	createConcatFile(concatFile, videos)

	err := ffmpeg_go.Input(concatFile, ffmpeg_go.KwArgs{"f": "concat", "safe": "0"}).
		Output(output, ffmpeg_go.KwArgs{"c": "copy"}).
		Run()
	if err != nil {
		log.Fatalf("Error running ffmpeg: %v", err)
	}
	return nil
}

func createConcatFile(filename string, videos []string) error {
	content := ""
	for _, video := range videos {
		content += fmt.Sprintf("file '%s'\n", video)
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("could not write concat file: %w", err)
	}

	return nil
}
