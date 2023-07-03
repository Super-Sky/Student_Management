package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
)

var config C
var configPath = "./config"

//Init 根据环境变量读取并加载相应配置文件
func Init(appEnv string, options ...func(*C)) C {
	if err := Read(appEnv, &config); err != nil {
		logrus.WithError(err).Warn("Fail to load config file")
	}
	for _, option := range options {
		option(&config)
	}

	return config
}

//Read 读取config
func Read(env string, config interface{}) error {
	in := "config"
	if env != "" {
		in += "." + env
	}
	viper.SetConfigName(in)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}
	if env != "" {
		f, err := os.Open(filepath.Join(configPath, in+".yml"))
		if err != nil {
			return fmt.Errorf("fatal error config file: %s", err)
		}
		defer f.Close()
		viper.MergeConfig(f)
	}
	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}
	return nil
}

//Config 获取config
func Config() C {
	return config
}

//Set 修改配置项
func Set(name string, value interface{}) {
	it := reflect.ValueOf(&config).Elem()
	val := it.FieldByName(name)
	v := reflect.ValueOf(value)
	val.Set(v)
}

//C 配置文件结构
type C struct {
	GinLogLevel string // gin的日志级别：release / debug / test
	RpcServer   struct {
		BaseUrl string // rpc url
		Port    string // rpc服务端口
	}
	HttpServer struct {
		BaseUrl string // http url
		Port    string // http服务端口
	}
	Log struct {
		Level    string // 日志级别：debug 、info、 warn（默认）  、error
		Path     string // 日志文件夹路径
		ToFile   string // 是否输出日志到文件
		ToScreen string // 是否输出日志到标准输出
	}
	Database struct {
		DbLogMode    bool   // 是否开启gorm日志
		LogZap       string // gorm日志级别："silent", "Silent"  ｜  "error", "Error"  ｜ "warn", "Warn" ｜ "info", "Info" ｜ "zap", "Zap"
		DbHost       string
		DbPort       string
		Dbname       string
		DbUsername   string
		DbPassword   string
		DbConfig     string
		MaxIdleConns int
		MaxOpenConns int
	}
	RedisBase struct {
		RedisAddr     string
		RedisPassword string
		Db            int
		MasterName    string
		RedisType     string // redis为本机redis模式，sentinel为redis集群哨兵模式
		MaxIdle       int    // redisPool数量
	}
	SnowServerAddr     string // 雪花算法服务地址
	PlatformServerAddr string // platform服务地址
	GithubProxy        string // github代理
	System             struct {
		IsSkipVerify int    // 是否跳过证书验证(只支持gitlab) 1: 是, 其他: 否
		OpenDir      string // 服务间公开文件地址
		IsDelFile    bool   // 是否删除文件
		ScriptPath   string // 脚本存放目录
		TimeoutFlag  int    // 处理超时逻辑
		Storage      string // 结果数据存储方式
	}
	TencentCos struct {
		Bucket    string
		Region    string
		SecretID  string
		SecretKey string
	}
}
