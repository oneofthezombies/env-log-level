package loglevelenv_test

import (
	"os"
	"testing"

	"github.com/oneofthezombies/loglevelenv"
)

func TestNoTargetGoLogNotExist(t *testing.T) {
	_, ok := os.LookupEnv("GO_LOG")
	if ok {
		t.Fail()
	}
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogEmpty(t *testing.T) {
	os.Setenv("GO_LOG", "")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogDebug(t *testing.T) {
	os.Setenv("GO_LOG", "debug")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestNoTargetGoLogInfo(t *testing.T) {
	os.Setenv("GO_LOG", "info")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogWarn(t *testing.T) {
	os.Setenv("GO_LOG", "warn")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "warn" {
		t.Fail()
	}
}

func TestNoTargetGoLogError(t *testing.T) {
	os.Setenv("GO_LOG", "error")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "error" {
		t.Fail()
	}
}

func TestNoTargetGoLogDebugUpper(t *testing.T) {
	os.Setenv("GO_LOG", "DEBUG")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "DEBUG" {
		t.Fail()
	}
}

func TestNoTargetGoLogInfoUpper(t *testing.T) {
	os.Setenv("GO_LOG", "INFO")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "INFO" {
		t.Fail()
	}
}

func TestNoTargetGoLogWarnUpper(t *testing.T) {
	os.Setenv("GO_LOG", "WARN")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "WARN" {
		t.Fail()
	}
}

func TestNoTargetGoLogErrorUpper(t *testing.T) {
	os.Setenv("GO_LOG", "ERROR")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "ERROR" {
		t.Fail()
	}
}

func TestNoTargetGoLogFoo(t *testing.T) {
	os.Setenv("GO_LOG", "foo")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "foo" {
		t.Fail()
	}
}

func TestNoTargetGoLogComma(t *testing.T) {
	os.Setenv("GO_LOG", ",")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogEqual(t *testing.T) {
	os.Setenv("GO_LOG", "=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogCommaEqual(t *testing.T) {
	os.Setenv("GO_LOG", ",=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogFooEmpty(t *testing.T) {
	os.Setenv("GO_LOG", "foo=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogTargetFooDebug(t *testing.T) {
	os.Setenv("GO_LOG", "foo=debug")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogTargetFooEqual(t *testing.T) {
	os.Setenv("GO_LOG", "foo==")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogTargetFooComma(t *testing.T) {
	os.Setenv("GO_LOG", "foo=,")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetGoLogTargetFooBar(t *testing.T) {
	os.Setenv("GO_LOG", "foo=bar")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooGoLogFooEmpty(t *testing.T) {
	os.Setenv("GO_LOG", "foo=")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooGoLogFooDebug(t *testing.T) {
	os.Setenv("GO_LOG", "foo=debug")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestTargetFooGoLogFooEqual(t *testing.T) {
	os.Setenv("GO_LOG", "foo==")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooGoLogFooComma(t *testing.T) {
	os.Setenv("GO_LOG", "foo=,")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooGoLogTargetFooBar(t *testing.T) {
	os.Setenv("GO_LOG", "foo=bar")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "bar" {
		t.Fail()
	}
}

func TestTargetFooGoLogDebug(t *testing.T) {
	os.Setenv("GO_LOG", "debug")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestTargetFooGoLogDebugTargetFooBar(t *testing.T) {
	os.Setenv("GO_LOG", "debug,foo=bar")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "bar" {
		t.Fail()
	}
}

func TestDefaultDebug(t *testing.T) {
	level, err := loglevelenv.Config().Default("debug").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestDefaultFoo(t *testing.T) {
	level, err := loglevelenv.Config().Default("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "foo" {
		t.Fail()
	}
}

func TestEnvFoo(t *testing.T) {
	os.Setenv("foo", "bar")
	level, err := loglevelenv.Config().Env("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "bar" {
		t.Fail()
	}
}

func TestMultipleTarget(t *testing.T) {
	os.Setenv("GO_LOG", "debug,foo=error,bar=warn")
	{
		level, err := loglevelenv.Parse()
		if err != nil {
			t.Error(err)
		}
		if level != "debug" {
			t.Fail()
		}
	}
	{
		level, err := loglevelenv.Config().Target("foo").Parse()
		if err != nil {
			t.Error(err)
		}
		if level != "error" {
			t.Fail()
		}
	}
	{
		level, err := loglevelenv.Config().Target("bar").Parse()
		if err != nil {
			t.Error(err)
		}
		if level != "warn" {
			t.Fail()
		}
	}
}
