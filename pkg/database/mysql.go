package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"log"
	"os"
	"time"
)

func MySQL(user, password, host, port, name string) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", user, password, host, port, name)
	timezone := "Asia/Seoul"
	_, err := time.LoadLocation(timezone)
	if err != nil {
		log.Panicf("Error loading location: %v", err.Error())
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
		log.Panicf("Error opening connection: %v", err.Error())
	}
	log.Println("GET INSTANCE DONE")
	return db
}
