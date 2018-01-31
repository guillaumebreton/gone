package util

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/guillaumebreton/gone/painter"
	"github.com/guillaumebreton/gone/state"
)

// Timer count time and update the state accordingly.
type Timer struct {
	state    *state.State
	command  string
	ticker   *time.Ticker
	painter  *painter.Painter
	notifier Notifier
}

// NewTimer create a new timer using a state, a command to execute
// and a painter to draw the screen.
func NewTimer(s *state.State, p *painter.Painter, c string, n Notifier) *Timer {
	return &Timer{
		state:    s,
		painter:  p,
		command:  c,
		notifier: n,
	}
}

// Run launch a timer and write the counter using the writer.
func (t *Timer) Run() {
	//start a new timer
	t.ticker = time.NewTicker(250 * time.Millisecond)
	i := 1
	for _ = range t.ticker.C {
		t.painter.Draw()
		if i > 4 && t.state.IsRunning() {
			i = 1
			t.state.Decrease()
			if t.state.IsEnded() {
				t.notifier.Notify("Pomodoro", t.state.StatusMessage())
				break
			}
			continue
		}
		i++
	}
	t.ticker.Stop()
	if t.command != "" {
		v := strings.Split(t.command, " ")
		cmd := exec.Command(v[0], v[1:]...)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Fail to execute command : %s - %v\n", t.command, err)
		}
	}
	t.state.WaitForConfirm(t.Run)
	t.painter.Draw()
}

// Stop the timer.
func (t *Timer) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
}
