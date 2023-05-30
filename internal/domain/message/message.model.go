package message

import (
	"time"
)

type Message struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	SenderID   uint      `gorm:"not null" comment:"보내는 사람" json:"sender_id,omitempty"`
	ReceiverID uint      `gorm:"not null" comment:"받는 사람" json:"receiver_id,omitempty"`
	Content    string    `gorm:"not null" comment:"내용" json:"content,omitempty"`
	CreatedAt  time.Time `comment:"생성일자" json:"created_at"`
	UpdatedAt  time.Time `comment:"수정일자" json:"updated_at"`
}

func (m *Message) TableName() string {
	return "message"
}