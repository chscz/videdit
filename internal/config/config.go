package config

import "encoding/json"

type Config struct {
	MariaDB MariaDB
	Video   Video
}

func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

func LoadConfig() (Config, error) {
	cfg := Config{
		MariaDB: MariaDB{
			UserName: "root",
			Password: "1111",
			Host:     "localhost",
			Port:     "3306",
			Schema:   "videdit",
		},
		Video: Video{
			UploadFilePath: "./upload",
			OutputFilePath: "./output",
		},
	}

	return cfg, nil
}
