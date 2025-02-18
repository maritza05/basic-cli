package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printVersion() {
	fmt.Printf("Version: %s\nCommit: %s\nDate: %s", version, commit, date)
}

func main() {
	versionFlag := flag.Bool("version", false, "print version information and other stuff")
	flag.Parse()

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}
	fmt.Println("That's all!!")
	fmt.Println("And an extra line")
}
