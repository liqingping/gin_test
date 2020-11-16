package logging

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/rifflock/lfshook"
)

type DefaultFieldHook struct {}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	return nil
}

func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func N()(f *DefaultFieldHook){
	return  &DefaultFieldHook{}
}


func FsHook(maxRemainCnt uint) logrus.Hook {
	//日志文件夹及文件名
	logName := "./logs/log"
	writerInfo, err := rotatelogs.New(
		logName+".%Y%m%d",

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),

		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithRotationCount(maxRemainCnt),
	)
	writerError, err := rotatelogs.New(
		logName+"-error.%Y%m%d",
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	//log.SetLevel(logrus.ErrorLevel)
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writerInfo,
		logrus.InfoLevel:  writerInfo,
		logrus.WarnLevel:  writerInfo,
		logrus.ErrorLevel: writerError,
		logrus.FatalLevel: writerError,
		logrus.PanicLevel: writerError,
	}, &logrus.JSONFormatter{})

	return lfsHook
}