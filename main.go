package main

import (
	"log"

	"github.com/chscz/videdit/internal/config"
	"github.com/chscz/videdit/internal/ffmpeg"
	"github.com/chscz/videdit/internal/handler"
	"github.com/chscz/videdit/internal/mariadb"
	"github.com/chscz/videdit/internal/router"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)

	// init database
	db, err := mariadb.InitMariaDB(cfg.MariaDB)
	if err != nil {
		log.Fatal(err)
	}
	repo := mariadb.Repository{DB: db}

	// init handler
	editor := ffmpeg.NewVideoEditor(cfg.Video)
	vh := handler.NewVideoHandler(repo, editor, cfg.Video)

	// set router
	r := router.InitRouter(vh)

	// server start
	log.Fatal(r.Start(":3000"))
}
