# Battery

[![Travis CI](https://travis-ci.org/jessfraz/battery.svg?branch=master)](https://travis-ci.org/jessfraz/battery)

Battery status getter written in Go.

## Installation

#### Binaries

- **linux** [386](https://github.com/jessfraz/battery/releases/download/v0.0.0/battery-linux-386) / [amd64](https://github.com/jessfraz/battery/releases/download/v0.0.0/battery-linux-amd64) / [arm](https://github.com/jessfraz/battery/releases/download/v0.0.0/battery-linux-arm) / [arm64](https://github.com/jessfraz/battery/releases/download/v0.0.0/battery-linux-arm64)

#### Via Go

```bash
$ go get github.com/jessfraz/battery
```

## Usage

```console
$ battery -h
 _           _   _
| |__   __ _| |_| |_ ___ _ __ _   _
| '_ \ / _` | __| __/ _ \ '__| | | |
| |_) | (_| | |_| ||  __/ |  | |_| |
|_.__/ \__,_|\__|\__\___|_|   \__, |
                              |___/

 Linux battery status checker.
 Version: v0.0.0
 Build: 00bf69e

  -d    run in debug mode
  -name string
        name of your battery (default "BAT0")
  -v    print version and exit (shorthand)
  -version
        print version and exit
```
