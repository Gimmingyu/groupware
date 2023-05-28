package database

type Model interface {
	TableName() string
	FindAll(limit, offset int) ([]Model, error)
	FindOne(id uint) (Model, error)
}
