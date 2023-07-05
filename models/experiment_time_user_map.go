package models

const (
	T_ExperimentTimeUserMap = "experiment_time_user_map"
)

// ExperimentTimeUserMap [...]
type ExperimentTimeUserMap struct {
	ID               int    `gorm:"primaryKey;column:id" json:"-"`
	UserID           int    `gorm:"column:user_id" json:"user_id"`
	ExperimentTimeID int    `gorm:"column:experiment_time_id" json:"experiment_time_id"`
	AScore           int    `gorm:"column:a_score" json:"a_score"` // 预习分数
	BScore           int    `gorm:"column:b_score" json:"b_score"` // 操作分数
	CScore           int    `gorm:"column:c_score" json:"c_score"` // 总结分数
	Score            int    `gorm:"column:score" json:"score"`     // 总分
	Desc             string `gorm:"column:desc" json:"desc"`       // 备注(占比信息)
}

func (e *ExperimentTimeUserMap) GetTableName() string {
	return T_ExperimentTimeUserMap
}