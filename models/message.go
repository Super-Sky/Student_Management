package models

import "time"

const (
	T_Message = "message"
)

// Message [...]
type Message struct {
	ID        int       `gorm:"primaryKey;column:id" json:"-"`
	TypeID    int       `gorm:"column:type_id" json:"type_id"` // 类型id
	Title     string    `gorm:"column:title" json:"title"`
	EnTitle   string    `gorm:"column:en_title" json:"en_title"`
	Source    string    `gorm:"column:source" json:"source"`         // 来源
	Info      string    `gorm:"column:info" json:"info"`             // 摘要
	IsSubmit  int       `gorm:"column:is_submit" json:"is_submit"`   // 是否发布(1-发布  2-不发布)
	Content   string    `gorm:"column:content" json:"content"`       // 文本内容
	AnnexPath string    `gorm:"column:annex_path" json:"annex_path"` // 附件地址
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func (m *Message) GetTableName() string {
	return T_Message
}
