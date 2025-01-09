package logrus_test

import (
	"os"
	"testing"

	"github.com/oneofthezombies/loglevelenv"
)

func TestMultipleTarget(t *testing.T) {
	os.Setenv("LOGLEVEL", "debug,foo=error,bar=warn")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
}
