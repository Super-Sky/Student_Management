package models

const (
	T_Role = "role"
)

// Role [...]
type Role struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"`
	Name   string `gorm:"column:name" json:"name"`       // 角色名称
	Info   string `gorm:"column:info" json:"info"`       // 角色描述
	EnName string `gorm:"column:en_name" json:"en_name"` // 英文名称
}

func (r *Role) GetTableName() string {
	return T_Role
}