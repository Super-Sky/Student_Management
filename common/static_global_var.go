package common

/**
 * 全局变量
 */
var (
	APP_ENV   string //pro-生产 、 test-测试 、  dev或空-开发
	SAAS_RUN  string //true是saas环境 ，否则是私有化
	ROOT_PATH string //项目启动位置
)

const (
	// 角色
	// 管理员
	RoleAdmin = "admin"
	// 教师
	RoleTeacher = "teacher"
	// 学生
	RoleStudent = "student"
)
