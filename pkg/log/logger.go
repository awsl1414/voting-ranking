/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-08-18 17:01:27
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-08-18 17:17:57
 * @FilePath: /voting-ranking/pkg/log/logger.go
 * @Description:
 *
 */

// log 日志

package log

import (
	"os"
	"path/filepath"
	"time"
	"voting-ranking/common/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger       // 全局 logrus 日志对象（用于控制台输出）
var logToFile *logrus.Logger // 全局 logrus 日志对象（用于文件输出）

// 日志文件名
var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

// 初始化
func init() {
	// 设置日志文件路径和文件名
	setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
}

// Log 方法，根据配置返回日志对象
func Log() *logrus.Logger {
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		// 控制台输出
		if log == nil {
			// 初始化 logrus 对象并设置格式和输出目标
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2008-01-0115:04:05"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

// logFile 方法，初始化并返回用于文件输出的日志对象
func logFile() *logrus.Logger {
	if logToFile == nil {
		// 初始化 logrus 对象并设置日志级别
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)
		// 创建一个新的日志轮转对象 logWriter，返回写日志对象logWriter
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名称
			loggerFile+"_%Y%m%d.log",
			// 设置最大保存时间
			rotatelogs.WithMaxAge(30*24*time.Hour),
			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		//设置时间格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		// 新增 Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}
