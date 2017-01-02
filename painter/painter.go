package painter

import (
	"fmt"
	"github.com/guillaumebreton/gone/state"
	"github.com/nsf/termbox-go"
	"strings"
	"sync"
)

// ColorMode defines the color scheme of a painter
type ColorMode struct {
	Bg      termbox.Attribute
	TimerFG termbox.Attribute
	TextFG  termbox.Attribute
}

// Painter draw a timer using termbox
type Painter struct {
	refreshMutex *sync.Mutex
	state        *state.State
	mode         ColorMode
	debug        bool
}

// NewPainter create a new painter based on a state, the color mode and the debug mode
func NewPainter(state *state.State, m string, debug bool) *Painter {
	var mode ColorMode
	if m == "light" {
		mode = ColorMode{
			Bg:      termbox.ColorWhite,
			TimerFG: termbox.ColorRed,
			TextFG:  termbox.ColorBlack,
		}
	} else {
		mode = ColorMode{
			Bg:      termbox.ColorBlack,
			TimerFG: termbox.ColorRed,
			TextFG:  termbox.ColorWhite,
		}
	}
	return &Painter{
		state:        state,
		mode:         mode,
		debug:        debug,
		refreshMutex: &sync.Mutex{},
	}
}

// Init the painter by initialising termbox
func (p *Painter) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

// Close the painter
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

// draw the timer and the bottom message
// if the debug mode is enabled also draw the debug mode
func (p *Painter) Draw() {
	p.refreshMutex.Lock()
	s := p.state.Duration()
	termbox.Clear(p.mode.Bg, p.mode.Bg)
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
	p.drawMessage(px, py+timerHeight+1, bannerWidth, p.state.Message())
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

// drawMessage a message using termbox
func (p *Painter) drawMessage(x, y, w int, message string) {
	padLeft := (w - len(message)) / 2
	m := strings.ToUpper(message)
	for i := 0; i < padLeft; i++ {
		m = " " + m
	}
	ux := x
	for _, c := range m {
		termbox.SetCell(ux, y, c, p.mode.TextFG, p.mode.Bg)
		ux++
	}
}

// drawTimer draw the timer duration using termbox
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
				termbox.SetCell(ux, uy, ' ', p.mode.TimerFG, p.mode.TimerFG)
			}
			ux++
		}
		uy++
	}
	return ux, uy
}
