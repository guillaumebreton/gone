package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
	"sync"
)

type Painter struct {
	refreshMutex *sync.Mutex
	state        *State
	mode         string
	debug        bool
}

func NewPainter(state *State, mode string, debug bool) *Painter {
	return &Painter{
		state:        state,
		mode:         mode,
		debug:        debug,
		refreshMutex: &sync.Mutex{},
	}
}

func (p *Painter) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (p *Painter) Close() {
	termbox.Close()
}

func (p *Painter) hline(y int) {
	w, _ := termbox.Size()
	for x := 0; x < w; x += 2 {
		termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorRed)
		termbox.SetCell(x+1, y, ' ', termbox.ColorWhite, termbox.ColorWhite)
	}

}
func (p *Painter) vline(x int) {
	_, h := termbox.Size()
	for y := 0; y < h; y += 2 {
		termbox.SetCell(x, y, ' ', termbox.ColorWhite, termbox.ColorRed)
		termbox.SetCell(x, y+1, ' ', termbox.ColorWhite, termbox.ColorWhite)
	}

}

// draw the timer
func (p *Painter) draw() {
	p.refreshMutex.Lock()
	s := p.state.Duration()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()
	bannerWidth := p.width(s)
	timerHeight := 5
	bannerHeight := 5
	px := w/2 - bannerWidth/2
	py := h/2 - bannerHeight/2
	if p.debug {
		p.hline(h / 2)
		p.vline(w / 2)
	}
	p.drawTimer(px, py, s)
	p.drawMessage(px, py+timerHeight+1, bannerWidth, state.Message())
	termbox.Flush()
	p.refreshMutex.Unlock()
}

// Compute the length of the timer
func (p *Painter) width(s string) int {
	result := 0
	for _, c := range s {
		switch c {
		case ':':
			result += 2
		default:
			result += 7
		}
	}
	return result - 1
}

func (p *Painter) drawMessage(x, y, w int, message string) {
	padLeft := (w - len(message)) / 2
	m := strings.ToUpper(message)
	for i := 0; i < padLeft; i++ {
		m = " " + m
	}
	ux := x
	for _, c := range m {
		termbox.SetCell(ux, y, c, termbox.ColorWhite, termbox.ColorBlack)
		ux++
	}
}

func (p *Painter) drawTimer(x, y int, str string) (ux int, uy int) {
	ux = x
	uy = y
	for _, c := range str {
		ux, uy = p.drawChar(ux, y, c)
	}
	return ux, uy
}

// drawChar draw a char and return the updated x and y
func (p *Painter) drawChar(x, y int, c rune) (ux int, uy int) {
	uy = y
	v := smallFont[c]
	if v == nil {
		panic(fmt.Errorf("Char not found in font"))
	}
	for _, l := range v {
		ux = x
		for _, c := range l {
			if c == '#' {
				termbox.SetCell(ux, uy, ' ', termbox.ColorRed, termbox.ColorRed)
			}
			ux++
		}
		uy++
	}
	return ux, uy
}
