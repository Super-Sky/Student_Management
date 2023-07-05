package models

import "time"

const (
	T_Role = "role"
)

// Role [...]
type Role struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`             // 角色名称
	Info      string    `gorm:"column:info" json:"info"`             // 角色描述
	EnName    string    `gorm:"column:en_name" json:"en_name"`       // 英文名称
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (r *Role) GetTableName() string {
	return T_Role
}
