package models

import "time"

const (
	T_CourseTime = "course_time"
)

// CourseTime [...]
type CourseTime struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"` // 开始时间(按照此排序)
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`     // 结束时间
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (c *CourseTime) GetTableName() string {
	return T_CourseTime
}
