package log

import "os"
import "github.com/sirupsen/logrus"

var Logger = &logrus.Logger{
	Out:       os.Stdout,
	Level:     logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{},
}
