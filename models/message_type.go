package models

import "time"

const (
	T_MessageType = "message_type"
)

// MessageType [...]
type MessageType struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`
	EnName    string    `gorm:"column:en_name" json:"en_name"`
	Info      string    `gorm:"column:info" json:"info"`             // 备注
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (m *MessageType) GetTableName() string {
	return T_MessageType
}
