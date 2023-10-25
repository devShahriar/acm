package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// credit: https://www.mountedthoughts.com/golang-logger-interface/

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	//encoderConfig.TimeKey = "time" // This will change the key from ts to time
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case CInfo:
		return zapcore.InfoLevel
	case CWarn:
		return zapcore.WarnLevel
	case CDebug:
		return zapcore.DebugLevel
	case CError:
		return zapcore.ErrorLevel
	case CFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	cores := []zapcore.Core{}

	if config.EnableConsole {
		level := getZapLevel(config.ConsoleLevel)
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(getEncoder(config.ConsoleJSONFormat), writer, level)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	// AddCallerSkip skips 2 number of callers, this is important else the file that gets
	// logged will always be the wrapped file. In our case zap.go
	logger := zap.New(combinedCore,
		zap.AddCallerSkip(2),
		zap.AddCaller(),
	).Sugar()

	return &zapLogger{
		sugaredLogger: logger,
	}, nil
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugaredLogger.Debugf(format, args...)
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.sugaredLogger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugaredLogger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugaredLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugaredLogger.Fatalf(format, args...)
}

func (l *zapLogger) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *zapLogger) Debug(format string) {
	l.sugaredLogger.Debugf(format)
}

func (l *zapLogger) Info(format string) {
	l.sugaredLogger.Infof(format)
}

func (l *zapLogger) Warn(format string) {
	l.sugaredLogger.Warnf(format)
}

func (l *zapLogger) Error(format string) {
	l.sugaredLogger.Errorf(format)
}

func (l *zapLogger) Fatal(format string) {
	l.sugaredLogger.Fatalf(format)
}

func (l *zapLogger) Panic(format string) {
	l.sugaredLogger.Fatalf(format)
}

func (l *zapLogger) Print(format string) {
	fmt.Print(format)
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

func (l *zapLogger) WithField(key string, value interface{}) Logger {
	return l.WithFields(Fields{key: value})
}

