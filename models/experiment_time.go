package models

import "time"

const (
	T_ExperimentTime = "experiment_time"
)

// ExperimentTime [...]
type ExperimentTime struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`
	ExperimentID   int       `gorm:"column:experiment_id" json:"experiment_id"`     // 实验id
	CourseTimeID   int       `gorm:"column:course_time_id" json:"course_time_id"`   // 学期课程时间id
	ExperimentDate time.Time `gorm:"column:experiment_date" json:"experiment_date"` // 实验时间(日期)
	CreatUserID    int       `gorm:"column:creat_user_id" json:"creat_user_id"`     // 创建者
	Status         string    `gorm:"column:status" json:"status"`
	UserMax        int       `gorm:"column:user_max" json:"user_max"`           // 最大人数
	UserMin        int       `gorm:"column:user_min" json:"user_min"`           // 最小开课人数
	ScoresStatus   int       `gorm:"column:scores_status" json:"scores_status"` // 成绩状态(1-已提交 2-未提交)
	ARatio         int       `gorm:"column:a_ratio" json:"a_ratio"`             // 预习占比
	BRatio         int       `gorm:"column:b_ratio" json:"b_ratio"`             // 操作占比
	CRatio         int       `gorm:"column:c_ratio" json:"c_ratio"`             // 总结占比
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`       // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`       // 更新时间
}

func (e *ExperimentTime) GetTableName() string {
	return T_ExperimentTime
}
