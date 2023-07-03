package common

import (
	"context"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	GVA_LOG *Logger
)

func InitLog(l *Logger) {
	GVA_LOG = l
}

var (
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

type Options struct {
	LogFileDir    string //文件保存地方
	AppName       string //日志文件前缀
	ErrorFileName string
	WarnFileName  string
	InfoFileName  string
	DebugFileName string
	Level         zapcore.Level //日志等级
	MaxSize       int           //日志文件小大（M）
	MaxBackups    int           // 最多存在多少个切片文件
	MaxAge        int           //保存的最大天数
	Development   bool          //是否是开发模式
	zap.Config
}

type ModOptions func(options *Options)

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts      *Options `json:"opts"`
	zapConfig zap.Config
	inited    bool
	lumError  *lumberjack.Logger
	lumWarn   *lumberjack.Logger
	lumInfo   *lumberjack.Logger
	lumDebug  *lumberjack.Logger
	TaskId    string
}

func NewLogger(taskID string, mod ...ModOptions) *Logger {
	l := &Logger{}
	l.Lock()
	defer l.Unlock()
	if l.inited {
		l.Info("[NewLogger] logger Inited")
		return nil
	}
	logLevel := Config().Log.Level
	var ll zapcore.Level
	switch logLevel {
	case "debug":
		ll = zapcore.DebugLevel
	case "info":
		ll = zapcore.InfoLevel
	case "warn":
		ll = zapcore.WarnLevel
	case "error":
		ll = zapcore.ErrorLevel
	default:
		ll = zapcore.WarnLevel
	}
	l.Opts = &Options{
		LogFileDir:    Config().Log.Path,
		AppName:       "app_log",
		ErrorFileName: "error.log",
		WarnFileName:  "warn.log",
		InfoFileName:  "info.log",
		DebugFileName: "debug.log",
		Level:         ll,
		MaxSize:       500,
		MaxBackups:    60,
		MaxAge:        30,
		Development:   true,
	}
	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "log" + sp + time.Now().Format("20060102") + sp
	}
	l.zapConfig = zap.NewDevelopmentConfig()
	l.zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}
	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}
	for _, fn := range mod {
		fn(l.Opts)
	}
	fmt.Println("l.Opts.Level:", l.Opts.Level)
	l.zapConfig.Level.SetLevel(l.Opts.Level)
	l.init(taskID)
	l.inited = true
	l.Info("[NewLogger] success")
	l.Logger = l.Logger.WithOptions(zap.AddCallerSkip(1))
	return l
}

func (l *Logger) init(taskId string) {
	l.setSyncers(taskId)
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	//consoleEncoder := zapcore.NewConsoleEncoder(l.zapConfig.EncoderConfig)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	if APP_ENV == "" {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	} else {
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && zapcore.ErrorLevel-l.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		zapcore.NewCore(fileEncoder, warnWS, warnPriority),
		zapcore.NewCore(fileEncoder, infoWS, infoPriority),
		zapcore.NewCore(fileEncoder, debugWS, debugPriority),
	}
	if l.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func (l *Logger) setSyncers(taskId string) {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + time.Now().Format("20060102") + sp + l.Opts.AppName + "-" + fN,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}
	if len(taskId) > 0 {
		l.TaskId = taskId
		errWS = l.zapCoreErrorWrite(l.Opts.ErrorFileName, taskId)
		warnWS = l.zapCoreErrorWrite(l.Opts.WarnFileName, taskId)
		infoWS = l.zapCoreErrorWrite(l.Opts.InfoFileName, taskId)
		debugWS = l.zapCoreErrorWrite(l.Opts.DebugFileName, taskId)
	} else {
		errWS = f(l.Opts.ErrorFileName)
		warnWS = f(l.Opts.WarnFileName)
		infoWS = f(l.Opts.InfoFileName)
		debugWS = f(l.Opts.DebugFileName)
	}
	return
}

func (l *Logger) zapCoreErrorWrite(fN, t string) zapcore.WriteSyncer {
	lum := &lumberjack.Logger{
		Filename:   l.Opts.LogFileDir + sp + time.Now().Format("20060102") + sp + t + sp + l.Opts.AppName + "-" + fN,
		MaxSize:    1024,
		MaxBackups: l.Opts.MaxBackups,
		MaxAge:     l.Opts.MaxAge,
		Compress:   true,
		LocalTime:  true,
	}
	switch fN {
	case l.Opts.ErrorFileName:
		l.lumError = lum
	case l.Opts.WarnFileName:
		l.lumWarn = lum
	case l.Opts.InfoFileName:
		l.lumInfo = lum
	case l.Opts.DebugFileName:
		l.lumDebug = lum
	}
	return zapcore.AddSync(lum)
}

func (l *Logger) close() {
	l.Logger.Sync()
	l.lumError.Close()
	l.lumWarn.Close()
	l.lumInfo.Close()
	l.lumDebug.Close()
}

func Debug(ctx context.Context, msg string, attributes ...zap.Field) {
	ctxLogger, ok := ctx.Value("logger").(*Logger)
	if ok {
		ctxLogger.Logger.With(zap.String("link_id", ctx.Value("link_id").(string))).Debug(msg, attributes...)
	} else {
		GVA_LOG.With(zap.String("link_id", ctx.Value("link_id").(string))).Debug(msg, attributes...)
	}
}

// Warn record warn
func Warn(ctx context.Context, msg string, attributes ...zap.Field) {
	ctxLogger, ok := ctx.Value("logger").(*Logger)
	if ok {
		ctxLogger.Logger.With(zap.String("link_id", ctx.Value("link_id").(string))).Warn(msg, attributes...)
	} else {
		GVA_LOG.With(zap.String("link_id", ctx.Value("link_id").(string))).Warn(msg, attributes...)
	}
}

// Info record info
func Info(ctx context.Context, msg string, attributes ...zap.Field) {
	ctxLogger, ok := ctx.Value("logger").(*Logger)
	if ok {
		ctxLogger.Logger.With(zap.String("link_id", ctx.Value("link_id").(string))).Info(msg, attributes...)
	} else {
		var id = ""
		if _, ok = ctx.Value("link_id").(string); ok {
			id = ctx.Value("link_id").(string)
		}
		GVA_LOG.With(zap.String("link_id", id)).Info(msg, attributes...)
	}
}

// Error record error
func Error(ctx context.Context, msg string, attributes ...zap.Field) {
	ctxLogger, ok := ctx.Value("logger").(*Logger)
	if ok {
		ctxLogger.Logger.With(zap.String("link_id", ctx.Value("link_id").(string))).Error(msg, attributes...)
	} else {
		GVA_LOG.With(zap.String("link_id", ctx.Value("link_id").(string))).Error(msg, attributes...)
	}
}

func LoggerClose(ctx context.Context) {
	ctxLogger, ok := ctx.Value("logger").(*Logger)
	if ok {
		if ctxLogger == nil {
		} else {
			ctxLogger.close()
		}
	}
}

func SetContextLogger(ctx context.Context, logger *Logger) context.Context {
	ctx = context.WithValue(ctx, "link_id", logger.TaskId)
	ctx = context.WithValue(ctx, "logger", logger)
	return ctx
}

func SetContext(ctx context.Context, tractId string) context.Context {
	ctx = context.WithValue(ctx, "link_id", tractId)
	return ctx
}

func GetFileWithLineNum() string {
	_, file, _, _ := runtime.Caller(0)
	// compatible solution to get gorm source directory with various operating systems
	gormSourceDir := regexp.MustCompile(`utils.utils\.go`).ReplaceAllString(file, "")
	// the second caller usually from gorm internal, so set i start from 2
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, gormSourceDir) || strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
