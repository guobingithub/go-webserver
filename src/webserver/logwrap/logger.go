package logwrap

import (
	"os"

	"apusic/go-webserver/src/config"
	"github.com/sirupsen/logrus"
)

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

var levelMap = map[string]logrus.Level{
	DebugLevel: logrus.DebugLevel,
	InfoLevel:  logrus.InfoLevel,
	WarnLevel:  logrus.WarnLevel,
	ErrorLevel: logrus.ErrorLevel,
	FatalLevel: logrus.FatalLevel,
}

func parseLevel(levelName string) logrus.Level {
	for k, v := range levelMap {
		if k == levelName {
			return v
		}
	}

	return logrus.InfoLevel // default level
}

type StructureLogger struct {
	*logrus.Logger
	MustSetFields []string
}

func NewStructureLogger(cfg config.LoggerConfig) *StructureLogger {
	slogger := &StructureLogger{logrus.New(), cfg.MustSetFields}
	slogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	slogger.SetOutput(os.Stdout)
	slogger.SetLevel(parseLevel(cfg.Level))
	slogger.SetReportCaller(cfg.EnableContext)
	slogger.MustSetFields = cfg.MustSetFields

	return slogger
}

func (slogger *StructureLogger) checkFields(fields logrus.Fields) *StructureLogger {
	for _, key := range slogger.MustSetFields {
		if _, ok := fields[key]; !ok {
			slogger.Fatalf("log entry don't have must set field: %v", key)
		}
	}

	return slogger
}

func (slogger *StructureLogger) LogEntryWithFields(fields logrus.Fields) *logrus.Entry {
	return slogger.checkFields(fields).WithFields(fields)
}
