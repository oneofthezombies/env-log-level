package main

import (
	"fmt"

	"github.com/oneofthezombies/log-env-filter/a"
)

func main() {
	b, c := a.GetModuleName()
	fmt.Printf("[%s][%s]", b, c)
}
