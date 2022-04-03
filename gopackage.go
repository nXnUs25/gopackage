package gopackage

import "fmt"

const VERSION = "1.0.1"
const ver = VERSION

func Version() string {
	return ver
}

func ShowVersion() {
	fmt.Printf("Current version: %s\n", ver)
}
