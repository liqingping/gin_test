package logging

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

var log = logrus.New()
func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.AddHook(FsHook(30))
	log.SetLevel(logrus.TraceLevel)

	// 设置日志格式为json格式
	//log.SetFormatter(&log.JSONFormatter{})

	//log.AddHook(N())
	//log.AddHook(FsHook(30))
	//// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	log.SetOutput(os.Stdout)
}
func setPrefix(start, end int) string {
	depth := start
	msg :=""
	for{
		f, file, line, _ := runtime.Caller(depth)
		depth ++
		if depth < end {
			msg += file+":"+strconv.Itoa(line)+"-"+runtime.FuncForPC(f).Name()+";"
		} else {
			break
		}
	}
	return msg
}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		// 请求前
		c.Next()
		// 请求后
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		comment := c.Errors.String()

		log.Printf("%v | %3d | %13v | %15s | %-7s %s %s",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
			comment)
	}
}

func Error(v ...interface{}){
	source := setPrefix(2,4)
	entry := log.WithFields(logrus.Fields{"source": source})
	entry.Error(v)
}

func Panic(v ...interface{}){
	source := setPrefix(6, 9)
	path := v[1]
	entry := log.WithFields(logrus.Fields{"source": source, "path": path})
	entry.Errorf("%+v",v[0])
}