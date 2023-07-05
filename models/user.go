package models

import (
	"context"
	"go.uber.org/zap"
	"student/common"
	"time"
)

const (
	T_User = "user"
)

// User [...]
type User struct {
	ID             int       `gorm:"primaryKey;column:id" json:"id"`
	Account        string    `gorm:"column:account" json:"account"`                 // 账号
	Password       string    `gorm:"column:password" json:"password"`               // 密码(默认密码123456)
	Name           string    `gorm:"column:name" json:"name"`                       // 姓名
	Sex            string    `gorm:"column:sex" json:"sex"`                         // 性别(men | women)
	RoleID         int       `gorm:"column:role_id" json:"role_id"`                 // 角色id
	OrganizationID int       `gorm:"column:organization_id" json:"organization_id"` // 组织id(班級id)
	Phone          int       `gorm:"column:phone" json:"phone"`                     // 手机号
	Telephone      int       `gorm:"column:telephone" json:"telephone"`             // 固话
	Email          string    `gorm:"column:email" json:"email"`                     // 邮箱
	BirthDate      time.Time `gorm:"column:birth_date" json:"birth_date"`           // 出生日期
	EntryDate      time.Time `gorm:"column:entry_date" json:"entry_date"`           // 入职日期
	Addr           string    `gorm:"column:addr" json:"addr"`                       // 办公地址
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`           // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`           // 更新时间
	IDCard         string    `gorm:"column:id_card" json:"id_card"`                 // 身份证号
	QQ             int       `gorm:"column:qq" json:"qq"`                           // qq号
	IconPath       string    `gorm:"column:icon_path" json:"icon_path"`             // 头像地址
}

func (u *User) GetTableName() string {
	return T_User
}

func (u *User) GetUserList(ctx context.Context) ([]User, error) {
	var (
		err    error
		result = make([]User, 0)
	)
	query := common.DB().Table(u.GetTableName())
	query.Statement.Context = common.SetContext(query.Statement.Context, ctx.Value("link_id").(string))
	if err = query.Find(&result).Error; err != nil {
		common.Error(ctx, "获取用户列表失败", zap.Error(err))
		return result, err
	}
	return result, err
}
