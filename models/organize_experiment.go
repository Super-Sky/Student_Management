package models

import "time"

const (
	T_OrganizeExperimentMap = "organize_experiment_map"
)

// OrganizeExperimentMap [...]
type OrganizeExperimentMap struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`
	OrganizeID   int       `gorm:"column:organize_id" json:"organize_id"`     // 组织id
	CourseID     int       `gorm:"column:course_id" json:"course_id"`         // 课程id
	ExperimentID int       `gorm:"column:experiment_id" json:"experiment_id"` // 实验id
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`       // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`       // 更新时间
}

func (o *OrganizeExperimentMap) GetTableName() string {
	return T_OrganizeExperimentMap
}
