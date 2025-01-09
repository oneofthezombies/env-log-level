package tests

import (
	"runtime/debug"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestIntegration(t *testing.T) {
	// GO_LOG=info
	// GO_LOG=debug,foo=info
	// GO_LOG=foo=debug
	// GO_LOG=[target][=][level][,...]
	// log.SetLevel(log.DebugLevel)
	info, ok := debug.ReadBuildInfo()
	moduleName := ""
	if ok {
		moduleName = info.Main.Path
	}
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Debug("A walrus appears")
	t.Errorf("hello [%s]", moduleName)
}
