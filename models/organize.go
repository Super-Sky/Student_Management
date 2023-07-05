package models

import "time"

const (
	T_Organize = "organize"
)

// Organize [...]
type Organize struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	Pid       int       `gorm:"column:pid" json:"pid"`
	Name      string    `gorm:"column:name" json:"name"`             // 组织名称
	UserCount int       `gorm:"column:user_count" json:"user_count"` // 用户数量
	Phone     int       `gorm:"column:phone" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (o *Organize) GetTableName() string {
	return T_Organize
}