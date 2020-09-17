# Battery

[![make-all](https://github.com/jessfraz/battery/workflows/make%20all/badge.svg)](https://github.com/jessfraz/battery/actions?query=workflow%3A%22make+all%22)
[![make-image](https://github.com/jessfraz/battery/workflows/make%20image/badge.svg)](https://github.com/jessfraz/battery/actions?query=workflow%3A%22make+image%22)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/battery)

Battery status getter written in Go.

**Table of Contents**

<!-- toc -->

- [Installation](#installation)
    + [Binaries](#binaries)
    + [Via Go](#via-go)
- [Usage](#usage)

<!-- tocstop -->

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/battery/releases).

#### Via Go

```console
$ go get github.com/jessfraz/battery
```

## Usage

```console
$ battery -h
battery -  Linux battery status checker.

Usage: battery <command>

Flags:

  -d      enable debug logging (default: false)
  --name  name of your battery (default: BAT0)

Commands:

  version  Show the version information.
```