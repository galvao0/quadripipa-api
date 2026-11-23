package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "> DEBUG: ", logger.Flags()),
		info:    log.New(writer, "> INFO: ", logger.Flags()),
		warning: log.New(writer, "> WARNING: ", logger.Flags()),
		err:     log.New(writer, "> ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println()
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println()
}
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println()
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println()
}

func (l *Logger) Debugf(f string, v ...interface{}) {
	l.debug.Printf(f, v...)
}
func (l *Logger) Infof(f string, v ...interface{}) {
	l.info.Printf(f, v...)
}
func (l *Logger) Warnf(f string, v ...interface{}) {
	l.warning.Printf(f, v...)
}
func (l *Logger) Errorf(f string, v ...interface{}) {
	l.err.Printf(f, v...)
}
