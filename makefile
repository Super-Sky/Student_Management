# -------------------------------------------------------------
# This makefile defines the following targets
#
#   - build (default) - build  binary for linux os
#   - server - run  server
#   - check - runs all code checks
# -------------------------------------------------------------


# TODO 以下项编译时需要根据实际情况进行修正
APP_VERSION             := v0.0.0
APP_COMPANY             := maxt

# windows环境时间串有一些差异处理
ifeq ($(OS),Windows_NT)
APP_BUILD_COMMIT            := $(shell git rev-parse --short HEAD)
APP_BUILD_BRANCH            := $(shell git symbolic-ref --short -q HEAD)
APP_BUILD_BUILD_TIME        := $(shell echo %Date:~0,4%%Date:~5,2%%Date:~8,2%)
else
APP_BUILD_COMMIT            := $(shell  git rev-parse --short HEAD)
APP_BUILD_BRANCH            := $(shell  git symbolic-ref --short -q HEAD)
APP_BUILD_BUILD_TIME        := $(shell  date "+%F %T")
endif

# 附加参数 用于指定是否输出gdb调试信息
BUILD_EXTRA		        :=

# windows平台命令兼容
ifeq ($(OS),Windows_NT)
    RM_CMD := del
    ENV_EXPORT_CMD :=set
else
    RM_CMD := rm
    ENV_EXPORT_CMD :=export
endif

build_saas:
	$(ENV_EXPORT_CMD) GOOS=linux  &&  go build  -ldflags \
    	" \
    	${BUILD_EXTRA}    \
    	-X 'gitee.com/super_sky/mkh_utils.CompanyLogo=${APP_COMPANY}'      \
        -X 'gitee.com/super_sky/mkh_utils.Version=${APP_VERSION}'     \
        -X 'gitee.com/super_sky/mkh_utils.Branch=${APP_BUILD_BRANCH}'     \
        -X 'gitee.com/super_sky/mkh_utils.Commit=${APP_BUILD_COMMIT}'    \
        -X 'gitee.com/super_sky/mkh_utils.BuildTime=${APP_BUILD_BUILD_TIME}'   \
    	" \
    	-o  student  main.go


# 本地开发启动服务
server:
	$(ENV_EXPORT_CMD) GOOS=""
	go run . api-server