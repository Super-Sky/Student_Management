package models

import "time"

const (
	T_Lab = "lab"
)

// Lab [...]
type Lab struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`
	Num       int       `gorm:"column:num" json:"num"`               // 容纳人数
	AdminID   int       `gorm:"column:admin_id" json:"admin_id"`     // 管理员id
	Addr      string    `gorm:"column:addr" json:"addr"`             // 实验室地点
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (l *Lab) GetTableName() string {
	return T_Lab
}
