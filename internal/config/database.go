package config

import (
	"gorm.io/gorm"
	"groupware/pkg/database"
	"os"
)

func Database(AppMode string) *gorm.DB {
	switch AppMode {
	case "production":
		return nil // mysql
	case "development":
		return database.Sqlite(os.Getenv("DB_NAME"))
	default:
		return nil // tunnelled mysql
	}
}
