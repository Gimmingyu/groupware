package user

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	Email     string    `gorm:"not null" comment:"이메일" json:"email,omitempty"`
	Password  string    `gorm:"not null" comment:"패스워드" json:"password,omitempty"`
	CreatedAt time.Time `comment:"생성일자" json:"created_at"`
	UpdatedAt time.Time `comment:"수정일자" json:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}