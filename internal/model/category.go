package model

type Category struct {
	Id       uint       `gorm:"primaryKey" json:"id,omitempty"`
	Text     string     `gorm:"not null" json:"text"`
	ParentID *uint      `json:"parent_id,omitempty"`
	Parent   *Category  `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Children []Category `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
