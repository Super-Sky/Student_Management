package v1

import (
	"context"
	"go.uber.org/zap"
	"student/common"
	"student/models"
)

type User struct{}

func (User) GetUserList(ctx context.Context, intput UserListRequest) ([]models.User, error) {
	var (
		err    error
		result []models.User
	)
	userModel := &models.User{}
	if result, err = userModel.GetUserList(ctx); err != nil {
		common.Error(ctx, "获取用户列表失败", zap.Error(err))
		return result, err
	}
	return result, err
}
