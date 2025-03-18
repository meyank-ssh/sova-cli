package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Fatal
)

var levelNames = map[LogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
	Fatal:   "FATAL",
}

type Logger struct {
	level  LogLevel
	output io.Writer
	prefix string
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		output: os.Stderr,
		prefix: "",
	}
}

func NewLoggerWithPrefix(level LogLevel, prefix string) *Logger {
	return &Logger{
		level:  level,
		output: os.Stderr,
		prefix: prefix,
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) Log(level LogLevel, format string, args ...interface{}) {
	if level >= l.level {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		prefix := ""
		if l.prefix != "" {
			prefix = fmt.Sprintf("[%s] ", l.prefix)
		}

		message := fmt.Sprintf(format, args...)
		logLine := fmt.Sprintf("%s %s[%s] %s\n", timestamp, prefix, levelNames[level], message)

		switch level {
		case Debug:
			fmt.Fprint(l.output, logLine)
		case Info:
			color.New(color.FgBlue).Fprint(l.output, logLine)
		case Warning:
			color.New(color.FgYellow).Fprint(l.output, logLine)
		case Error, Fatal:
			color.New(color.FgRed).Fprint(l.output, logLine)
		}
	}

	if level == Fatal {
		os.Exit(1)
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.Log(Debug, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.Log(Info, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.Log(Warning, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.Log(Error, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.Log(Fatal, format, args...)
}

var DefaultLogger = NewLogger(Info)
