package logs

import "github.com/natefinch/lumberjack"

func newAccessLogWriter() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "/var/log/app/access.log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	}
}

func newErrorLogWriter() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "/var/log/app/error.log",
		MaxSize:    10, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	}
}
