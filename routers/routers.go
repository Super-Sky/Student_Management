package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"net/http"
	"student/common"
	"time"
)

// InitRouter 初始化路由
func InitRouter(c common.C) *gin.Engine {
	gin.SetMode(c.GinLogLevel)
	r := gin.New()
	r.Use(gin.Recovery())
	// 跨域处理
	r.Use(Cors())
	//心跳
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/test", func(c *gin.Context) {
		RedisAndMysqlTest(c)
	})
	return r
}

func RedisAndMysqlTest(c *gin.Context) {
	var (
		client redis.Conn
		err    error
		s      = time.Now().String()
		s2     string
	)
	if client, err = common.GetRedis(); err != nil {
		common.GVA_LOG.Error("get redis conn fail", zap.Error(err))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if _, err = client.Do("SET", "ping", s); err != nil {
		common.GVA_LOG.Error("set to redis fail", zap.Error(err))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if s2, err = redis.String(client.Do("GET", "ping")); err != nil {
		common.GVA_LOG.Error("get from redis  fail", zap.Error(err))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	common.GVA_LOG.Info("test redis:", zap.String("data", s2))
	c.String(http.StatusOK, "ok")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("starttime", time.Now().UnixNano()/int64(time.Millisecond))
		method := c.Request.Method
		//fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, developerId, Cache-Control,CustomerCode")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,CustomerCode")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.Header("Cache-Control", "private, max-age=86400")
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
		return
	}
}
