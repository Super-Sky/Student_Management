package common

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	_db *gorm.DB
)

//DB 获取DB
func DB() *gorm.DB {
	if _db == nil {
		panic("please init db first")
	}
	return _db
}

//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func InitGorm(c C) error {
	var (
		err error
	)
	_db, err = GormMysql(c)
	return err
}

//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql(c C) (*gorm.DB, error) {
	ctx := SetContextLogger(context.Background(), GVA_LOG)
	m := c.Database
	if m.Dbname == "" {
		Error(ctx, "MySQL用户名配置为空")
		panic(errors.New("MySQL用户名配置为空"))
	}
	dsn := m.DbUsername + ":" + m.DbPassword + "@tcp(" + m.DbHost + ":" + m.DbPort + ")/" + m.Dbname + "?" + m.DbConfig
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.DbLogMode)); err != nil {
		Error(ctx, "MySQL启动异常", zap.Error(err))
		panic(err)
	} else {
		Info(ctx, "[NewGorm] success")
		sqlDB, err := db.DB()
		if err != nil {
			Error(ctx, "db.DB err:"+err.Error())
			panic(err)
		}
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db, nil
	}
}

//@description: 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	c := Config()
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch c.Database.LogZap {
	case "silent", "Silent":
		config.Logger = Default.LogMode(gormLogger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(gormLogger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(gormLogger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(gormLogger.Info)
	case "zap", "Zap":
		config.Logger = Default.LogMode(gormLogger.Info)
	default:
		if mod {
			config.Logger = Default.LogMode(gormLogger.Info)
			break
		}
		config.Logger = Default.LogMode(gormLogger.Silent)
	}
	return config
}
