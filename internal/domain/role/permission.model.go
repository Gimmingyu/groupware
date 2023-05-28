package role

type Permission struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	RoleID uint   `gorm:"not null" comment:"역할은 여러가지 권한을 가진다" json:"role_id,omitempty"`
	Name   string `gorm:"not null" comment:"권한 이름" json:"name,omitempty"`
}

func (p *Permission) TableName() string {
	//TODO implement me
	panic("implement me")
}

func (p *Permission) FindAll(limit, offset int) ([]*Permission, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Permission) FindOne(id uint) (*Permission, error) {
	//TODO implement me
	panic("implement me")
}
