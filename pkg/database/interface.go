package database

import "gorm.io/gorm"

type Model interface {
	TableName() string
	FindAll(db *gorm.DB, limit, offset int) ([]Model, error)
	FindOne(db *gorm.DB, id uint) (Model, error)
}
