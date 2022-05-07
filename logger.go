package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	runtimeLogger "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

type (
	CustomLogger interface {
		Info(format string, args ...interface{})
		Warning(format string, args ...interface{})
		Error(format string, args ...interface{})
		Debug(format string, args ...interface{})
		Trace(format string, args ...interface{})
	}

	customLogger struct {
		log   *logrus.Entry
		level int
	}

	PrefixLoggerName struct {
		Title       string
		Description string
	}
)

func NewLogger(setupName []PrefixLoggerName, level int) CustomLogger {
	formatter := &runtimeLogger.Formatter{ChildFormatter: &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}}
	formatter.Line = true
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stdout)
	// logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)

	if len(setupName) > 1 {
		for i := 0; i < len(setupName[1:]); i++ {
			logrus.WithField(setupName[i].Title, setupName[i].Description)
		}
	}

	return &customLogger{
		log:   logrus.WithField(setupName[0].Title, setupName[0].Description),
		level: level}
}

func (c *customLogger) Info(format string, args ...interface{}) {
	if format != "" {
		c.log.Infof(format, args)
	} else {
		c.log.Info(args)
	}
}

func (c *customLogger) Warning(format string, args ...interface{}) {
	c.log.Info(fileInfo)

	if format != "" {
		c.log.Warningf(format, args)
	} else {
		c.log.Warning(args)
	}
}

func (c *customLogger) Error(format string, args ...interface{}) {
	c.log.Info(fileInfo)

	if format != "" {
		c.log.Errorf(format, args)
	} else {
		c.log.Error(args)
	}
}

func (c *customLogger) Debug(format string, args ...interface{}) {
	c.log.Info(fileInfo)

	if format != "" {
		c.log.Debugf(format, args)
	} else {
		c.log.Debug(args)
	}
}

func (c *customLogger) Trace(format string, args ...interface{}) {
	c.log.Info(fileInfo)

	if format != "" {
		c.log.Tracef(format, args)
	} else {
		c.log.Trace(args)
	}
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
