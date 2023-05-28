package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func dsn(user, password, host, port string, name string) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", user, password, host, port, name)
}

func Sqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("group.db"), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // Disable color
			}),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
