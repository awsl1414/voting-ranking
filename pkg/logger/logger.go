/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-11 22:46:44
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-11 22:51:07
 * @FilePath: /hnuahe-presentation-voting-ranking/pkg/logger/logger.go
 * @Description:
 *
 */
package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	LOG_DIR string = "./runtime/log"
	mu      sync.Mutex
)

func init() {
	// 设置日志格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{

		// 作为时间格式的参考值，因此这个格式字符串表示日期和时间的格式。
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

// 写入一般日志消息到指定文件
func Write(msg string, filename string) {
	setOutputFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

// 写入调试级别的日志消息
func Debug(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args...)

}

// 写入信息级别的日志消息
func Info(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args...)
}

// 写入警告级别的日志消息
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args...)
}

// 写入严重错误级别的日志消息
func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args...)
}

// 写入错误级别的日志消息
func Error(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args...)
}

// 写入恐慌级别的日志消息
func Panic(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args...)
}

// 写入跟踪级别的日志消息
func Trace(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args...)
}

// 检查并创建日志目录
func checkLogDir() {
	if _, err := os.Stat(LOG_DIR); os.IsNotExist(err) { // 检查目录是否存在
		err = os.MkdirAll(LOG_DIR, 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir '%s' err: %s", LOG_DIR, err))
		}
	}
}

// 创建日志文件
func createLogFile(link string, logName string) *os.File {
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(LOG_DIR, logName+link+timeStr+".log")

	out, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("open log file err", err)
	}
	return out

}

// 设置日志输出文件和日志级别
func setOutputFile(level logrus.Level, logName string) {
	// 加互斥锁保证线程安全 setOutputFile和logrus.SetOutput线程不安全
	mu.Lock()
	defer mu.Unlock()

	checkLogDir() // 检查并创建日志目录

	out := createLogFile("-", logName)

	logrus.SetOutput(out)  // 设置日志输出
	logrus.SetLevel(level) // 设置日志级别

}

// 返回 gin 的日志配置，用于将日志输出到文件和标准输出
func LoggerToFile() gin.LoggerConfig {

	checkLogDir()

	out := createLogFile("success_", "")

	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, out), // 同时输出日志到文件和控制台
	}
	return conf

}

// gin 中间件，用于捕获 panic 并记录错误日志
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			checkLogDir()
			f := createLogFile("error_", "")
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic err time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})

			// 终止后续接口调用，不加的话，recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}

	}()

	c.Next()
}
