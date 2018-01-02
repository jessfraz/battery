package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jessfraz/battery/version"
	"github.com/sirupsen/logrus"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = ` _           _   _
| |__   __ _| |_| |_ ___ _ __ _   _
| '_ \ / _` + "`" + ` | __| __/ _ \ '__| | | |
| |_) | (_| | |_| ||  __/ |  | |_| |
|_.__/ \__,_|\__|\__\___|_|   \__, |
                              |___/

 Linux battery status checker.
 Version: %s
 Build: %s

`
)

var (
	name  string
	debug bool
	vrsn  bool
)

func init() {
	// parse flags
	flag.StringVar(&name, "name", "BAT0", "name of your battery")

	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION, version.GITCOMMIT))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("battery version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	battery, err := New(name)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := battery.GetStatus(); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infoln(battery.String())
}
