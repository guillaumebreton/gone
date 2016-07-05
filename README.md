# Gone

Gone is a simple pomodoro timer from cli for OSX and Linux

# Installation

see [release page](https://github.com/guillaumebreton/gone/releases) to get the
right artifact

# Usage
```
Usage of ./gone:
  -e string
        The command to execute when a session is done
  -l int
        Duration of a long break (default 15)
  -p string
        Pattern to apply (for example wswswl) (default "wswswl")
  -s int
        Duration of a short break (default 5)
  -w int
        Duration of a working session (default 25)
```

# Example

```
./gone -w 25 -l 30 -s 5 -e "say done"
```
