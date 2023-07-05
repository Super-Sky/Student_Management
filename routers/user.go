package routers

import (
	"github.com/gin-gonic/gin"
	"student/controller"
)

func User(r *gin.RouterGroup) {
	user := r.Group("/user")
	user.POST("/get_user_list", controller.UserController{}.GetUserList)
}
