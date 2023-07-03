package main

import (
	"context"
	"fmt"
	"gitee.com/super_sky/mkh_utils"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	_ "gorm.io/driver/mysql"
	"os"
	"student/common"
	"student/routers"
	"student/timer"
)

func main() {
	//输出版本信息
	mkh_utils.CheckVersion()
	common.APP_ENV = mkh_utils.AppEnv
	common.ROOT_PATH, _ = os.Getwd()

	//加载配置文件
	configer := common.Init(common.APP_ENV)

	// 初始化日志
	logger := common.NewLogger("")
	common.InitLog(logger)
	ctx := common.SetContextLogger(context.Background(), logger)
	common.Info(ctx, "日志级别:"+configer.Log.Level)
	common.Info(ctx, "环境变量", zap.Any("APP_ENV", common.APP_ENV), zap.Any("SAAS_RUN", common.SAAS_RUN))
	//todo 加载证书文件

	// 初始化数据库
	if configer.Database.DbHost == "" {
		common.Info(ctx, "skip  mysql init")
	} else {
		common.InitGorm(configer)
	}

	// 初始化redis
	if configer.RedisBase.RedisAddr == "" {
		common.Info(ctx, "skip  redis init")
	} else {
		common.InitRedis(configer)
		common.Info(ctx, "[NewRedigo] success")
	}
	timer.Init()
	common.Info(ctx, "初始化定时器", zap.Any("status", "成功"))
	common.Info(ctx, "app server port:", zap.Any("port", configer.HttpServer.Port))
	//设置gin路由模式
	app := cli.NewApp()
	app.Name = "task Application"
	port := configer.HttpServer.Port
	app.Commands = []cli.Command{
		{
			Name:  "api-server",
			Usage: "run api server,such as(./task api-server)",
			Action: func(cliContext *cli.Context) {
				engine := routers.InitRouter(configer)
				if err := engine.Run(":" + port); err != nil {
					common.Info(ctx, "api-server run err :"+err.Error(), zap.String("http服务启动端口", port))
					return
				}
			},
		},
	}

	//打印打包信息
	app.Commands = append(app.Commands, cli.Command{
		Name:  "version",
		Usage: "version info print ,such as(./task -version)",
		Action: func(cliContext *cli.Context) {
			common.Info(ctx, fmt.Sprintf("current branch:%s", mkh_utils.Branch))
			common.Info(ctx, fmt.Sprintf("current commit:%s", mkh_utils.Commit))
			common.Info(ctx, fmt.Sprintf("current build time:%s", mkh_utils.BuildTime))
		},
		After: func(context *cli.Context) error {
			common.Info(ctx, "mkh exec end")
			return nil
		},
	})

	//运行带参数的自定义命令
	app.Commands = append(app.Commands, cli.Command{
		Name:    "test",
		Aliases: nil,
		Usage:   "run custom command with args  ,such as(./task test --path ./static/jump_data.json --limit 2  )",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "path", Value: "./static/data.json", Usage: "file path"},
			cli.IntFlag{Name: "limit", Value: 10, Usage: "goroutine chan limit"},
		},
		Action: func(cliContext *cli.Context) {
			arg1 := cliContext.String("path")
			arg2 := cliContext.Int("limit")
			common.Info(ctx, fmt.Sprintf("path:%s, limit:%d", arg1, arg2))

		},
		After: func(context *cli.Context) error {
			common.Info(ctx, "test exec end")
			return nil
		},
	})
	common.Info(ctx, "[NewServer] success", zap.String("http服务启动端口", port))
	if err := app.Run(os.Args); err != nil {
		common.Info(ctx, "app run err :"+err.Error())
	}
}
