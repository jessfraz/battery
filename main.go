package main

import (
	"flag"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

const (
	// VERSION is the binary version
	VERSION = "v0.1.0"
)

var (
	name    string
	debug   bool
	version bool
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&name, "name", "BAT0", "name of your battery")
	flag.Parse()
}

func main() {
	// set log level
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	if version {
		fmt.Println(VERSION)
		return
	}

	battery, err := New(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := battery.GetStatus(); err != nil {
		log.Fatal(err)
	}

	log.Infoln(battery.String())
}
