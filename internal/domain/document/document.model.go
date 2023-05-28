package document

import (
	"time"
)

type Document struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	GroupID   uint      `gorm:"not null" comment:"그룹 id" json:"group_id,omitempty"`
	AuthorID  uint      `gorm:"not null" comment:"작성자 id" json:"author_id,omitempty"`
	Email     string    `gorm:"not null" comment:"제목" json:"email,omitempty"`
	Password  string    `gorm:"not null" comment:"패스워드" json:"password,omitempty"`
	CreatedAt time.Time `comment:"생성일자" json:"created_at"`
	UpdatedAt time.Time `comment:"수정일자" json:"updated_at"`
}

func (d *Document) TableName() string {
	return "document"
}

func (d *Document) FindAll(limit, offset int) ([]*Document, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Document) FindOne(id uint) (*Document, error) {
	//TODO implement me
	panic("implement me")
}
