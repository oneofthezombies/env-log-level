package a

import (
	"runtime/debug"
)

func GetModuleName() (string, string) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", ""
	}
	return info.Path, info.Main.Path
}
