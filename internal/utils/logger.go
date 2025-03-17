package utils

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

// LogLevel represents the severity level of a log message
type LogLevel int

const (
	// Debug level for detailed information
	Debug LogLevel = iota
	// Info level for general information
	Info
	// Warning level for potential issues
	Warning
	// Error level for errors that don't stop the program
	Error
	// Fatal level for errors that stop the program
	Fatal
)

var levelNames = map[LogLevel]string{
	Debug:   "DEBUG",
	Info:    "INFO",
	Warning: "WARNING",
	Error:   "ERROR",
	Fatal:   "FATAL",
}

// Logger is a simple logging utility
type Logger struct {
	level  LogLevel
	output io.Writer
	prefix string
}

// NewLogger creates a new logger with the specified level
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		output: os.Stderr,
		prefix: "",
	}
}

// NewLoggerWithPrefix creates a new logger with the specified level and prefix
func NewLoggerWithPrefix(level LogLevel, prefix string) *Logger {
	return &Logger{
		level:  level,
		output: os.Stderr,
		prefix: prefix,
	}
}

// SetOutput sets the output writer for the logger
func (l *Logger) SetOutput(output io.Writer) {
	l.output = output
}

// SetLevel sets the log level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// SetPrefix sets the prefix for log messages
func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

// Log logs a message at the specified level
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

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	l.Log(Debug, format, args...)
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	l.Log(Info, format, args...)
}

// Warning logs a warning message
func (l *Logger) Warning(format string, args ...interface{}) {
	l.Log(Warning, format, args...)
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	l.Log(Error, format, args...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.Log(Fatal, format, args...)
}

// DefaultLogger is the default logger instance
var DefaultLogger = NewLogger(Info)
