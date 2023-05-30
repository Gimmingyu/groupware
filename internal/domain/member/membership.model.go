package member

import (
	"gorm.io/gorm"
	"time"
)

type Membership struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	UserID    uint           `gorm:"not null" comment:"유저 id" json:"user_id,omitempty"`
	GroupID   uint           `gorm:"not null" comment:"그룹 id" json:"group_id,omitempty"`
	CreatedAt time.Time      `comment:"멤버가 된 날짜" json:"created_at"`
	UpdatedAt time.Time      `comment:"수정일자" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" comment:"삭제일자" json:"deleted_at"`
}

func (m *Membership) TableName() string {
	return "membership"
}