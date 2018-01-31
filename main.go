package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/guillaumebreton/gone/painter"
	"github.com/guillaumebreton/gone/state"
	"github.com/guillaumebreton/gone/util"
	"github.com/nsf/termbox-go"
)

var w = flag.Int("w", 25, "Duration of a working session")
var s = flag.Int("s", 5, "Duration of a short break")
var l = flag.Int("l", 15, "Duration of a long break")
var p = flag.String("p", "wswswl", "Pattern to  follow (for example wswswl)")
var e = flag.String("e", "", "The command to execute when a session is done")
var m = flag.String("m", "dark", "Select the color mode (light or dark)")
var n = flag.Bool("n", false, "Enable desktop notifications")
var d = flag.Bool("debug", false, "Debug option for development purpose")

var wg sync.WaitGroup

var currentState *state.State
var currentPainter *painter.Painter
var currentTimer *util.Timer

func main() {
	flag.Parse()
	if *p == "" {
		fmt.Printf("Invalid pattern ''%s', should not be empty\n", *p)
		os.Exit(2)
	}
	for _, c := range *p {
		if c != 'w' && c != 'l' && c != 's' {
			fmt.Printf("Invalid pattern ''%s', should contain only w,s, or l\n", *p)
			os.Exit(2)
		}
	}

	notifier := util.NewNullNotifier()
	if *n {
		notifier = util.NewDesktopNotifier()
	}

	currentState = state.NewState(*p, *w, *s, *l)
	currentPainter = painter.NewPainter(currentState, *m, *d)
	currentPainter.Init()
	currentTimer = util.NewTimer(currentState, currentPainter, *e, notifier)
	go handleKeyEvent()
	go currentTimer.Run()
	wg.Add(1)
	wg.Wait()
	os.Exit(1)

}

// handleKeyEvent handles keys on event.
func handleKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				exit()
			}
			switch ev.Ch {
			case 'q':
				exit()
			case 'p':
				if currentState.IsRunning() {
					currentState.Pause()
				} else {
					currentState.Resume()
				}
				currentPainter.Draw()
			case 'y':
				if currentState.IsWaiting() {
					currentState.Resume()
				}
			case 'Y':
				if currentState.IsWaiting() {
					currentState.Resume()
				}
			default:
				if currentState.IsWaiting() {
					exit()
				}
			}
		case termbox.EventResize:
			currentPainter.Draw()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

// exit kill the timer and destroy the painter.
func exit() {
	currentTimer.Stop()
	currentPainter.Close()
	wg.Done()
}
