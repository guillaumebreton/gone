package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"sync"
)

var w = flag.Int("w", 25, "Duration of a working session")
var s = flag.Int("s", 5, "Duration of a short break")
var l = flag.Int("l", 15, "Duration of a long break")
var p = flag.String("p", "wswswl", "Pattern to  follow (for example wswswl)")
var e = flag.String("e", "", "The command to execute when a session is done")
var m = flag.String("mode", "dark", "Select the color mode")
var d = flag.Bool("debug", false, "Debug option for development purpose")

var wg sync.WaitGroup

var state *State
var painter *Painter
var timer *Timer

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
	state = NewState(*p, *w, *s, *l)
	painter = NewPainter(state, *m, *d)
	painter.Init()
	timer = NewTimer(state, painter, *e)
	go handleKeyEvent()
	go timer.run()
	wg.Add(1)
	wg.Wait()
	os.Exit(1)

}

// handleKeyEvent handles keys on event
func handleKeyEvent() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Ch {
			case 'q':
				exit()
			case 'p':
				if state.IsRunning() {
					state.Pause()
				} else {
					state.Resume()
				}
				// TODO remove this call
				painter.draw()
			case 'y':
				if state.IsWaiting() {
					state.Resume()
				}
			case 'Y':
				if state.IsWaiting() {
					state.Resume()
				}
			default:
				if state.IsWaiting() {
					exit()
				}
			}
		case termbox.EventResize:
			painter.draw()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func exit() {
	timer.Stop()
	painter.Close()
	wg.Done()
}
