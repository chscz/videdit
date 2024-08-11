package mariadb

import (
	"net"
	"time"

	"github.com/chscz/videdit/internal/config"
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMariaDB(cfg config.MariaDB) (*gorm.DB, error) {
	mysqlCfg := &mysql.Config{
		User:   cfg.UserName,
		Passwd: cfg.Password,
		Net:    "tcp",
		Addr:   net.JoinHostPort(cfg.Host, cfg.Port),
		DBName: cfg.Schema,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
		Loc:                  time.UTC,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dsn := mysqlCfg.FormatDSN()
	db, err := gorm.Open(gormmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
