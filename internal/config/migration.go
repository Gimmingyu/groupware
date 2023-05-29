package config

import (
	"gorm.io/gorm"
	"groupware/internal/domain/document"
	"groupware/internal/domain/group"
	"groupware/internal/domain/member"
	"groupware/internal/domain/message"
	"groupware/internal/domain/role"
	"groupware/internal/domain/user"
	"log"
)

func Migrate(db *gorm.DB) {

	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&user.User{}).TableName(), err)
	}
	if err := db.AutoMigrate(&document.Document{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&document.Document{}).TableName(), err)
	}
	if err := db.AutoMigrate(&group.Group{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&group.Group{}).TableName(), err)
	}
	if err := db.AutoMigrate(&group.Invitation{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&group.Invitation{}).TableName(), err)
	}
	if err := db.AutoMigrate(&member.Membership{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&member.Membership{}).TableName(), err)
	}
	if err := db.AutoMigrate(&message.Message{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&message.Message{}).TableName(), err)
	}
	if err := db.AutoMigrate(&role.Role{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&role.Role{}).TableName(), err)
	}
	if err := db.AutoMigrate(&role.Permission{}); err != nil {
		log.Fatalf("failed to migrate table %v : %v", (&role.Permission{}).TableName(), err)
	}
}
