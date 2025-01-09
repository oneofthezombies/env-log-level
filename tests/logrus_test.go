package tests

import (
	"runtime/debug"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestIntegration(t *testing.T) {
	// LOG_LEVEL=debug,foo=debug
	// log.SetLevel(log.DebugLevel)
	info, ok := debug.ReadBuildInfo()
	moduleName := ""
	if ok {
		moduleName = info.Main.Path
	}
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Debug("A walrus appears")
	t.Errorf("hello [%s]", moduleName)
}
