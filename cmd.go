package main

import (
	"os"
)

func isVerbose() bool {
	for _, a := range os.Args {
		if a == "-v" || a == "-verbose" {
			return true
		}
	}
	return false
}
