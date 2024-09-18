package logs

import (
	"os"
	"time"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/rs/zerolog"
)

var loggers struct {
	access zerolog.Logger
	err    zerolog.Logger
	info   zerolog.Logger
	warn   zerolog.Logger
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	cfg := getSysConfig()
	lv := getLogLevel(cfg)
	loggers.access = newAccess(cfg, lv)
	loggers.err = newError(cfg, lv)
	loggers.info = newInfo(lv)
	loggers.warn = newWarn(lv)
}

func getSysConfig() dconfig.Sys {
	if cfg, err := dconfig.New(); err == nil {
		return cfg.Sys
	}
	return dconfig.Sys{
		Env: dconfig.LOCAL,
		TZ:  "Asia/Tokyo",
	}
}

func getLogLevel(cfg dconfig.Sys) zerolog.Level {
	if cfg.IsProd() || cfg.IsStg() {
		return zerolog.InfoLevel
	}
	return zerolog.DebugLevel
}

func newAccess(cfg dconfig.Sys, lv zerolog.Level) zerolog.Logger {
	if cfg.IsLocal() {
		return newInfo(lv)
	}
	return zerolog.New(newAccessLogWriter()).
		Level(lv).
		With().
		Timestamp().
		Logger()
}

func newError(cfg dconfig.Sys, lv zerolog.Level) zerolog.Logger {
	if cfg.IsLocal() {
		return newInfo(lv)
	}
	return zerolog.New(newErrorLogWriter()).
		Level(lv).
		With().
		Timestamp().
		Logger()
}

func newInfo(lv zerolog.Level) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
	}).
		Level(lv).
		With().
		Timestamp().
		Logger()
}

func newWarn(lv zerolog.Level) zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
	}).
		Level(lv).
		With().
		Timestamp().
		Caller().
		Logger()
}
