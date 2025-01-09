package loglevelenv

import (
	"os"
	"strings"
)

type LogLevelEnvConfig struct {
	envKey       string
	target       string
	defaultValue string
}

func Config() *LogLevelEnvConfig {
	return &LogLevelEnvConfig{
		envKey:       "LOG_LEVEL",
		defaultValue: "info",
		target:       "",
	}
}

func (c *LogLevelEnvConfig) Env(envKey string) *LogLevelEnvConfig {
	c.envKey = envKey
	return c
}

func (c *LogLevelEnvConfig) Default(defaultValue string) *LogLevelEnvConfig {
	c.defaultValue = defaultValue
	return c
}

func (c *LogLevelEnvConfig) Target(target string) *LogLevelEnvConfig {
	c.target = target
	return c
}

func (c *LogLevelEnvConfig) Parse() (string, error) {
	envValue := os.Getenv(c.envKey)
	if envValue == "" {
		return c.defaultValue, nil
	}

	targetLevelPairs := strings.Split(envValue, ",")
	maybeLevel := c.defaultValue
	for _, targetLevelPair := range targetLevelPairs {
		targetLevel := strings.Split(targetLevelPair, "=")
		if len(targetLevel) == 1 {
			if targetLevel[0] != "" {
				maybeLevel = targetLevel[0]
			}
		} else if len(targetLevel) == 2 {
			target := targetLevel[0]
			if target == "" {
				continue
			}

			if target == c.target {
				if targetLevel[1] != "" {
					maybeLevel = targetLevel[1]
					break
				}
			}
		}
	}

	return maybeLevel, nil
}

func Parse() (string, error) {
	return Config().Parse()
}
