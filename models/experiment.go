package models

import "time"

const (
	T_Experiment = "experiment"
)

// Experiment [...]
type Experiment struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	LabID     int       `gorm:"column:lab_id" json:"lab_id"`       // 实验室
	CourseID  int       `gorm:"column:course_id" json:"course_id"` // 课程
	Name      string    `gorm:"column:name" json:"name"`
	Num       int       `gorm:"column:num" json:"num"`
	UserMax   int       `gorm:"column:user_max" json:"user_max"`     // 人数上限
	UserMin   string    `gorm:"column:user_min" json:"user_min"`     // 最少开课人数
	ClassTime int       `gorm:"column:class_time" json:"class_time"` // 课时
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (e *Experiment) GetTableName() string {
	return T_Experiment
}
