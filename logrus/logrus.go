package logrus

import (
	"os"

	"github.com/sirupsen/logrus"
)

func MapLevel(maybeLevel string) (logrus.Level, error) {
	debug := logrus.New()
	debug.SetOutput(os.Stderr)
	debug.SetLevel(logrus.DebugLevel)
	debug.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	go func() {
		debug.Debug("hello11")
	}()
	debug.Debug("hello11")
	return logrus.ParseLevel(maybeLevel)
}
