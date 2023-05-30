package group

import (
	"time"
)

type Invitation struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	UserID    uint      `gorm:"not null" comment:"링크 생성자" json:"user_id,omitempty"`
	GroupID   uint      `gorm:"not null" comment:"그룹 id" json:"group_id,omitempty"`
	Link      string    `gorm:"not null" comment:"링크" json:"link,omitempty"`
	ExpiredAt time.Time `comment:"만료일자" json:"expired_at"`
	CreatedAt time.Time `comment:"생성일자" json:"created_at"`
	UpdatedAt time.Time `comment:"수정일자" json:"updated_at"`
}

func (i *Invitation) TableName() string {
	return "invitation"
}