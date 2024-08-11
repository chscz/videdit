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

func InitRouter(wh *handler.WebHandler, vh *handler.VideoHandler) *echo.Echo {
	e := echo.New()

	// e.File("/", "public/index.html")
	e.Static("/", "public")
	e.POST("/upload", vh.UploadVideo)
	e.POST("/create_video", vh.CreateVideo)
	e.POST("/test", test)

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
	// 두 동영상 파일의 경로를 설정합니다.
	videoA := "./test/red.mp4"
	videoB := "./test/pink.mp4"
	videos := []string{videoA, videoB}
	output := "./test/output.mp4"

	// 두 동영상을 연속적으로 합치기 위해 임시 파일을 생성합니다.
	concatFile := "concat_list.txt"

	// concat 파일을 생성하여 비디오 파일 리스트를 작성합니다.
	createConcatFile(concatFile, videos)

	// FFmpeg를 사용하여 동영상 파일을 합칩니다.
	err := ffmpeg_go.Input(concatFile, ffmpeg_go.KwArgs{"f": "concat", "safe": "0"}).
		Output(output, ffmpeg_go.KwArgs{"c": "copy"}).
		Run()
	if err != nil {
		log.Fatalf("Error running ffmpeg: %v", err)
	}

	fmt.Println("동영상 파일을 성공적으로 합쳤습니다!")
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
