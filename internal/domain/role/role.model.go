package role

type Role struct {
	ID      uint   `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	GroupID uint   `gorm:"not null" comment:"그룹에서는 여러가지 역할을 정의할 수 있음" json:"group_id,omitempty"`
	Name    string `gorm:"not null" comment:"역할 이름" json:"name,omitempty"`
}

func (r *Role) TableName() string {
	//TODO implement me
	panic("implement me")
}

func (r *Role) FindAll(limit, offset int) ([]*Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Role) FindOne(id uint) (*Role, error) {
	//TODO implement me
	panic("implement me")
}
