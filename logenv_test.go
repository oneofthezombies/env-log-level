package logenv_test

import (
	"os"
	"testing"

	"github.com/oneofthezombies/logenv"
)

func TestParsing(t *testing.T) {
	os.Setenv("GO_LOG", "debug,foo=info,bar=a=b")
	logenv.LogEnv().Get()
}
