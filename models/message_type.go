package models

const (
	T_MessageType = "message_type"
)

// MessageType [...]
type MessageType struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"`
	Name   string `gorm:"column:name" json:"name"`
	EnName string `gorm:"column:en_name" json:"en_name"`
	Info   string `gorm:"column:info" json:"info"` // 备注
}

func (m *MessageType) GetTableName() string {
	return T_MessageType
}
