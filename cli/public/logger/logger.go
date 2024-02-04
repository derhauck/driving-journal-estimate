package logger

import (
	"fmt"
	"os"
)

type Level int

var levels = map[Level]string{
	DEBUG:   "DEBUG",
	INFO:    "INFO",
	WARNING: "WARNING",
	ERROR:   "ERROR",
	LOG:     "LOG",
}

func (l Level) String() string {
	return levels[l]
}

func ParseLevel(level string) (Level, error) {
	for k, v := range levels {
		if v == level {
			return k, nil
		}
	}
	return DEFAULT, fmt.Errorf("invalid log level '%s'", level)
}

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	LOG
	DEFAULT = WARNING
)

type logger struct {
	level         Level
	defaultWriter *os.File
}

func (l *logger) Info(v any) {
	l.bootstrapLogging(INFO, v)
}

func (l *logger) Infof(format string, v ...any) {
	l.bootstrapLogging(INFO, fmt.Sprintf(format, v...))
}

func (l *logger) Log(v any) {
	l.bootstrapLogging(LOG, v)
}

func (l *logger) Logf(format string, v ...any) {
	l.bootstrapLogging(LOG, fmt.Sprintf(format, v...))
}

func (l *logger) Warning(v any) {
	l.bootstrapLogging(WARNING, v)
}

func (l *logger) Warningf(format string, v ...any) {
	l.bootstrapLogging(WARNING, fmt.Sprintf(format, v...))
}

func (l *logger) Error(v any) {
	l.bootstrapLogging(ERROR, v)
}

func (l *logger) Errorf(format string, v ...any) {
	l.bootstrapLogging(ERROR, fmt.Sprintf(format, v...))
}

func (l *logger) bootstrapLogging(level Level, v any) bool {
	if l.level <= level {
		_, err := fmt.Fprint(l.getWriter(level), fmt.Sprintf("%s: %v\n", level, v))
		if err == nil {
			return true
		}
	}
	return false
}

func (l *logger) GetLevel() Level {
	return l.level
}

func (l *logger) SetLevel(level Level) {
	l.level = level
}

type Inf interface {
	Log(v any)
	Logf(format string, v ...any)
	Info(v any)
	Infof(format string, v ...any)
	Warning(v any)
	Warningf(format string, v ...any)
	Error(v any)
	Errorf(format string, v ...any)
	SetLevel(level Level)
	GetLevel() Level
}

func New(level Level) Inf {
	return &logger{
		level:         level,
		defaultWriter: os.Stdout,
	}
}

func (l *logger) getWriter(level Level) *os.File {
	return l.defaultWriter
}
