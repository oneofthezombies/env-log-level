package logrus_test

import (
	"testing"

	"github.com/oneofthezombies/loglevelenv/logrus"
)

func TestMultipleTarget(t *testing.T) {
	logrus.MapLevel("")
	logrus.MapLevel("")
	logrus.MapLevel("")
	logrus.MapLevel("")
	logrus.MapLevel("")
	logrus.MapLevel("")
	logrus.MapLevel("")
}
