package loglevelenv_test

import (
	"os"
	"testing"

	"github.com/oneofthezombies/loglevelenv"
)

func TestNoTargetLogLevelNotExist(t *testing.T) {
	_, ok := os.LookupEnv("LOG_LEVEL")
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

func TestNoTargetLogLevelEmpty(t *testing.T) {
	os.Setenv("LOG_LEVEL", "")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelDebug(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestNoTargetLogLevelInfo(t *testing.T) {
	os.Setenv("LOG_LEVEL", "info")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelWarn(t *testing.T) {
	os.Setenv("LOG_LEVEL", "warn")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "warn" {
		t.Fail()
	}
}

func TestNoTargetLogLevelError(t *testing.T) {
	os.Setenv("LOG_LEVEL", "error")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "error" {
		t.Fail()
	}
}

func TestNoTargetLogLevelDebugUpper(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "DEBUG" {
		t.Fail()
	}
}

func TestNoTargetLogLevelInfoUpper(t *testing.T) {
	os.Setenv("LOG_LEVEL", "INFO")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "INFO" {
		t.Fail()
	}
}

func TestNoTargetLogLevelWarnUpper(t *testing.T) {
	os.Setenv("LOG_LEVEL", "WARN")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "WARN" {
		t.Fail()
	}
}

func TestNoTargetLogLevelErrorUpper(t *testing.T) {
	os.Setenv("LOG_LEVEL", "ERROR")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "ERROR" {
		t.Fail()
	}
}

func TestNoTargetLogLevelFoo(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "foo" {
		t.Fail()
	}
}

func TestNoTargetLogLevelComma(t *testing.T) {
	os.Setenv("LOG_LEVEL", ",")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelEqual(t *testing.T) {
	os.Setenv("LOG_LEVEL", "=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelCommaEqual(t *testing.T) {
	os.Setenv("LOG_LEVEL", ",=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelFooEmpty(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelTargetFooDebug(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=debug")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelTargetFooEqual(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo==")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelTargetFooComma(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=,")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestNoTargetLogLevelTargetFooBar(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=bar")
	level, err := loglevelenv.Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooLogLevelFooEmpty(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooLogLevelFooDebug(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=debug")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestTargetFooLogLevelFooEqual(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo==")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooLogLevelFooComma(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=,")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "info" {
		t.Fail()
	}
}

func TestTargetFooLogLevelTargetFooBar(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foo=bar")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "bar" {
		t.Fail()
	}
}

func TestTargetFooLogLevelDebug(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	level, err := loglevelenv.Config().Target("foo").Parse()
	if err != nil {
		t.Error(err)
	}
	if level != "debug" {
		t.Fail()
	}
}

func TestTargetFooLogLevelDebugTargetFooBar(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug,foo=bar")
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
	os.Setenv("LOG_LEVEL", "debug,foo=error,bar=warn")
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
