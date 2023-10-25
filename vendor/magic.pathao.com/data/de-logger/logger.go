package logger

import (
	"errors"
	"fmt"
	"log"
	common "magic.pathao.com/data/esb_contract/golang/pathao/proto/v1/common"
	"sync"
)

// credit: https://www.mountedthoughts.com/golang-logger-interface/

// A global variable so that log functions can be directly accessed
var (
	once   sync.Once
	log1   Logger
)

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	CDebug = "debug"
	//Info is default log level
	CInfo = "info"
	//Warn is for logging messages about possible issues
	CWarn = "warn"
	//Error is for logging errors
	CError = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	CFatal = "fatal"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

const(
	serviceTypeKey = "srvType"
	cityKey        = "city"
	driverTypeKey  = "drvType"
	orderIdKey     = "orderId"
	traceIDKey     = "traceId"
)


var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Debug(format string)
	Info(format string)
	Warn(format string)
	Error(format string)
	Fatal(format string)
	Panic(format string)
	Print(format string)
	WithFields(keyValues Fields) Logger
	WithField(key string, value interface{}) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

// InitLogger logger for whole application
func init() {
	fmt.Println("I am from logger.go")
	once.Do(
		func() {

			config := Configuration{
				EnableConsole:     true,
				ConsoleLevel:      CInfo,
				ConsoleJSONFormat: true,
				EnableFile:        false,
			}
			err := newLogger(config, InstanceZapLogger)
			if err != nil {
				log.Fatalf("Could not instantiate log %s", err.Error())
			}

			//formatter1 := &log.TextFormatter{
			//	FullTimestamp:          true,
			//	TimestampFormat:        "2006-01-02T15:04:05.999999999Z07:00",
			//	DisableLevelTruncation: true,
			//	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//		return "", fmt.Sprintf(" %s:%d", f.Function, f.Line)
			//	},
			//}
			//logger = &log.Logger{
			//	Out:          os.Stdout,
			//	Level:        log.InfoLevel,
			//	Formatter:    formatter1,
			//	ReportCaller: true,
			//}

		})

}


//newLogger returns an instance of logger
func newLogger(config Configuration, loggerInstance int) error {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		log1 = logger
		log1.Info("zap logger is initialized")
		return nil

	case InstanceLogrusLogger:
		logger, err := newLogrusLogger(config)
		if err != nil {
			return err
		}
		log1 = logger
		log1.Info("logrus logger is initialized")
		return nil

	default:
		return errInvalidLoggerInstance
	}
}

func GetLogger(serviceType common.Enums_ServiceType, cityId int32, driverType common.Enums_DriverType, orderHash string) Logger {
	return WithFields(Fields{serviceTypeKey: serviceType, cityKey: cityId, driverTypeKey: driverType, orderIdKey: orderHash})
}

// GetSimpleLogger return logger using only traceID
func GetSimpleLogger(traceID string) Logger {
	return WithFields(Fields{traceIDKey: traceID})
}

//GetLoggerInstance get logger
func GetLoggerInstance() Logger {
	return log1
}


func Debugf(format string, args ...interface{}) {
	log1.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log1.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log1.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log1.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log1.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log1.Panicf(format, args...)
}

func Debug(format string) {
	log1.Debug(format)
}

func Info(format string) {
	log1.Info(format)
}

func Warn(format string) {
	log1.Warn(format)
}

func Error(format string) {
	log1.Error(format)
}

func Fatal(format string) {
	log1.Fatal(format)
}

func Panic(format string) {
	log1.Panic(format)
}

func WithFields(keyValues Fields) Logger {
	return log1.WithFields(keyValues)
}

func WithField(keyValues Fields) Logger {
	return log1.WithFields(keyValues)
}

