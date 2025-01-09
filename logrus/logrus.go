package logrus

import (
	"github.com/oneofthezombies/loglevelenv"
	"github.com/sirupsen/logrus"
)

func MapLevel(maybeLevel string) (logrus.Level, error) {
	return logrus.ParseLevel(maybeLevel)
}

func (c *loglevelenv.LogLevelEnvConfig) ParseLogrus() (logrus.Level, error) {

}
