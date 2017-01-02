# Gone [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Gone is a simple cli pomodoro timer for OSX and Linux. It can execute a
command every time a session is done. Responsiveness inside (lel)


![scrot](https://github.com/guillaumebreton/gone/raw/master/srot.png)


# Installation

see [release page](https://github.com/guillaumebreton/gone/releases) to get the
right artifact and put it in your path :)

# Usage
Run gone and Use ```q``` to quit, ```p``` to pause.

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
  -s int
        Duration of a short break (default 5)
  -w int
        Duration of a working session (default 25)
```

# Example

```
./gone -w 25 -l 30 -s 5 -e "say done"
```

# Release the application

Install https://github.com/goreleaser/releaser and execute :

```
  release
```
