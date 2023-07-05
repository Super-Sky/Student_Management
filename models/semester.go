package models

import "time"

const (
	T_Semester = "semester"
)

// Semester [...]
type Semester struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	Name      string    `gorm:"column:name" json:"name"`             // 学期名称
	StartTime time.Time `gorm:"column:start_time" json:"start_time"` // 开始时间
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`     // 结束时间
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (s *Semester) GetTableName() string {
	return T_Semester
}