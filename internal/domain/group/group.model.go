package group

type Group struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id,omitempty"`
	OwnerID     uint   `gorm:"not null" comment:"그룹 소유자" json:"owner_id,omitempty"`
	Name        string `gorm:"not null" comment:"그룹 이름" json:"name,omitempty"`
	Description string `comment:"긴 설명" json:"description,omitempty"`
	Summary     string `comment:"짧은 설명" json:"summary,omitempty"`
}

func (g *Group) TableName() string {
	return "group"
}