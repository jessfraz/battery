package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/genuinetools/pkg/cli"
	"github.com/jessfraz/battery/version"
	"github.com/sirupsen/logrus"
)

var (
	name  string
	debug bool
)

func main() {
	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "battery"
	p.Description = "Linux battery status checker"

	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.StringVar(&name, "name", "BAT0", "name of your battery")

	p.FlagSet.BoolVar(&debug, "d", false, "enable debug logging")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}

	// Set the main program action.
	p.Action = func(ctx context.Context, args []string) error {
		// On ^C, or SIGTERM handle exit.
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			for sig := range c {
				logrus.Infof("Received %s, exiting.", sig.String())
				os.Exit(0)
			}
		}()

		battery, err := New(name)
		if err != nil {
			logrus.Fatal(err)
		}

		if err := battery.GetStatus(); err != nil {
			logrus.Fatal(err)
		}

		logrus.Infoln(battery.String())

		return nil
	}

	// Run our program.
	p.Run()
}
