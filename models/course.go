package models

import "time"

const (
	T_Course = "course"
)

// Course [...]
type Course struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	CourseNum string    `gorm:"column:course_num" json:"course_num"` // 课程号
	Name      string    `gorm:"column:name" json:"name"`             // 课程名
	Ordinal   int       `gorm:"column:ordinal" json:"ordinal"`       // 课序
	ItemNum   int       `gorm:"column:item_num" json:"item_num"`     // 项数
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (c *Course) GetTableName() string {
	return T_Course
}