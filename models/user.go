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

type User struct {
	Id             int64     `json:"id"   gorm:"column:id; type:int"`                        // id
	Account        string    `json:"account"   gorm:"column:account; type:varchar(255)"`     // 账号
	Password       string    `json:"password"   gorm:"column:password; type:varchar(255)"`   // 密码
	Name           string    `json:"name"   gorm:"column:name; type:varchar(255)"`           // 姓名
	Sex            string    `json:"sex"   gorm:"column:sex; type:varchar(255)"`             // 性别
	RoleID         int64     `json:"role_id" gorm:"column:role_id;type:int"`                 // 角色id
	OrganizationID int64     `json:"organization_id" gorm:"column:organization_id;type:int"` //组织id
	Phone          int64     `json:"phone"   gorm:"column:phone; type:int"`                  // 电话号码
	Telephone      int64     `json:"telephone"   gorm:"column:telephone; type:int"`          // 固话
	Email          string    `json:"email"   gorm:"column:email; type:varchar(255)"`         // 邮箱
	BirthDate      time.Time `json:"birth_date" gorm:"column:birth_date; type:datetime"`     // 出生日期
	EntryDate      time.Time `json:"entry_date" gorm:"column:entry_date; type:datetime"`     // 入学日期
	Addr           string    `json:"addr" gorm:"column:addr; type:varchar(255)"`             // 办公
	IdCard         string    `json:"id_card" gorm:"column:id_card; type:varchar(255)"`       // 身份证
	QQ             string    `json:"qq" gorm:"column:qq; type:varchar(255)"`
	IconPath       string    `json:"icon_path" gorm:"column:icon_path; type:varchar(255)"` // 头像地址
	CreatedAt      time.Time `json:"created_at"   gorm:"column:created_at; type:datetime"` // 创建时间
	UpdatedAt      time.Time `json:"updated_at"   gorm:"column:updated_at; type:datetime"` // 更新时间
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
