package main

import (
	"fmt"
	"os"
)

var (
	Version    = "0"
	minversion string
)

func main() {
	info("goDCPU v%s", Version+minversion)
}

// info reports to stdout in exactly the same manner that fmt.Printf() behaves, except that it also appends a newline.
func info(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}

// error behaves exactly like info(), except that it reports to stderr instead of stdout.
func error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
}
