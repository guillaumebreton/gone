# Gone [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Gone is a simple cli pomodoro timer for OSX and Linux. It can execute a
command every time a session is done. Responsiveness inside (lel)


![scrot](https://github.com/guillaumebreton/gone/raw/master/srot.png)


# Installation

see [release page](https://github.com/guillaumebreton/gone/releases) to get the
right artifact and put it in your path :)

# Usage
Run gone. Press ```Esc``` to quit, ```p``` to pause. During a session
only ```Esc``` exits, so a stray key press will not stop the timer by
mistake. When a session ends, press ```y``` to continue to the next one.

```
Usage of ./bin/gone:
  -debug
        Debug option for development purpose
  -e string
        The command to execute when a session is done
  -l int
        Duration of a long break (default 15)
  -m string
        Select the color mode (default "dark")
  -p string
        Pattern to  follow (for example wswswl) (default "wswswl")
  -n    Enable desktop notifications
  -s int
        Duration of a short break (default 5)
  -sound string
        Path to a sound file played when a session ends
  -timer string
        Timer digit color: black, red, green, yellow, blue, magenta, cyan, white (default "red")
  -w int
        Duration of a working session (default 25)
```

# Example

```
./gone -w 25 -l 30 -s 5 -e "say done"
```

# Development

Gone requires a supported Go toolchain.

```
go test ./...
make build
```

# Release the application

Maintainers publish a release by pushing an annotated `vX.Y.Z` tag. GitHub
Actions builds the Darwin and Linux archives and publishes them with checksums
on the [release page](https://github.com/guillaumebreton/gone/releases).

```
git tag -a vX.Y.Z -m "vX.Y.Z"
git push origin vX.Y.Z
```
