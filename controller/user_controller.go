package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"student/common"
	"student/models"
	v1 "student/services/v1"
)

type UserController struct{}

func (UserController) GetUserList(c *gin.Context) {
	var (
		err     error
		input   v1.UserListRequest
		service v1.User
		result  []models.User
	)
	ctx := context.Background()
	ctx = common.SetContext(ctx, "GetUserList")
	if err = c.ShouldBind(&input); err != nil {
		common.Error(ctx, "DelHistory参数缺失", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "参数异常"})
		return
	}
	if result, err = service.GetUserList(ctx, input); err != nil {
		common.Error(ctx, "获取用户列表失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数异常"})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
