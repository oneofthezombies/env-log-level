package logenv

import (
	"fmt"
	"os"
	"strings"
)

type LogEnvBuilder struct {
	envKey       string
	target       string
	defaultValue string
}

func LogEnv() *LogEnvBuilder {
	return &LogEnvBuilder{
		envKey:       "GO_LOG",
		target:       "",
		defaultValue: "info",
	}
}

func (b *LogEnvBuilder) Env(envKey string) *LogEnvBuilder {
	b.envKey = envKey
	return b
}

func (b *LogEnvBuilder) Target(target string) *LogEnvBuilder {
	b.target = target
	return b
}

func (b *LogEnvBuilder) Default(defaultValue string) *LogEnvBuilder {
	b.defaultValue = defaultValue
	return b
}

func (b *LogEnvBuilder) Get() (string, error) {
	envValue := os.Getenv(b.envKey)
	if envValue == "" {
		return b.defaultValue, nil
	}
	targetLevelPairs := strings.Split(envValue, ",")
	for _, targetLevelPair := range targetLevelPairs {
		targetLevel := strings.SplitN(targetLevelPair, "=", 2)
		if len(targetLevel) == 1 {
			maybeLevel := targetLevel[0]
			println(maybeLevel)
		} else if len(targetLevel) == 2 {
			target := targetLevel[0]
			maybeLevel := targetLevel[1]
			println(target, maybeLevel)
		} else {
			return "", fmt.Errorf("%s is not <target>=<level> in %s", targetLevel, envValue)
		}
	}
	return "", nil
}
