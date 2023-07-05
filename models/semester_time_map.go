package models

const (
	T_SemesterTimeMap = "semester_time_map"
)

// SemesterTimeMap [...]
type SemesterTimeMap struct {
	ID         int `gorm:"primaryKey;column:id" json:"-"`
	TimeID     int `gorm:"column:time_id" json:"time_id"`         // 课程时间id
	SemesterID int `gorm:"column:semester_id" json:"semester_id"` // 学期id
}

func (s *SemesterTimeMap) GetTableName() string {
	return T_SemesterTimeMap
}