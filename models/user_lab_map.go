package models

import "time"

const (
	T_UserLabMap = "user_lab_map"
)

// UserLabMap [...]
type UserLabMap struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	LabID     int       `gorm:"column:lab_id" json:"lab_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (u *UserLabMap) GetTableName() string {
	return T_UserLabMap
}
