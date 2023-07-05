package models

const (
	T_UserLabMap = "user_lab_map"
)

// UserLabMap [...]
type UserLabMap struct {
	ID     int `gorm:"primaryKey;column:id" json:"-"`
	UserID int `gorm:"column:user_id" json:"user_id"`
	LabID  int `gorm:"column:lab_id" json:"lab_id"`
}

func (u *UserLabMap) GetTableName() string {
	return T_UserLabMap
}
