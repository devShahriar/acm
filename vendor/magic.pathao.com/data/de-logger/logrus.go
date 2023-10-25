package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// credit: https://www.mountedthoughts.com/golang-logger-interface/

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

func newLogrusLogger(config Configuration) (Logger, error) {
	lLogger, e := createLogrusLogger(config)
	if e != nil {
		return nil, e
	}

	return &logrusLogger{
		logger: lLogger,
	}, nil
}

func createLogrusLogger(config Configuration) (*logrus.Logger, error) {
	logLevel := config.ConsoleLevel
	if logLevel == "" {
		logLevel = config.FileLevel
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	stdOutHandler := os.Stdout
	lLogger := &logrus.Logger{
		Out:       stdOutHandler,
		Formatter: getFormatter(config.ConsoleJSONFormat),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}
	if config.EnableConsole && config.EnableFile {
		lLogger.SetOutput(io.Writer(stdOutHandler))
	}
	return lLogger, nil
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *logrusLogger) Debug(format string) {
	l.logger.Debug(format)
}

func (l *logrusLogger) Info(format string) {
	l.logger.Info(format)
}

func (l *logrusLogger) Warn(format string) {
	l.logger.Warn(format)
}

func (l *logrusLogger) Error(format string) {
	l.logger.Error(format)
}

func (l *logrusLogger) Fatal(format string) {
	l.logger.Fatal(format)
}

func (l *logrusLogger) Panic(format string) {
	l.logger.Fatal(format)
}

func (l *logrusLogger) Print(format string) {
	fmt.Print(format)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogger) WithField(key string, value interface{}) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(Fields{key: value})),
	}
}

func (l *logrusLogEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) Panicf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *logrusLogEntry) Debug(format string) {
	l.entry.Debug(format)
}

func (l *logrusLogEntry) Info(format string) {
	l.entry.Info(format)
}

func (l *logrusLogEntry) Warn(format string) {
	l.entry.Warn(format)
}

func (l *logrusLogEntry) Error(format string) {
	l.entry.Error(format)
}

func (l *logrusLogEntry) Fatal(format string) {
	l.entry.Fatal(format)
}

func (l *logrusLogEntry) Panic(format string) {
	l.entry.Fatalf(format)
}

func (l *logrusLogEntry) Print(format string) {
	fmt.Print(format)
}


func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) WithField(key string, value interface{}) Logger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(Fields{key: value})),
	}
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}

