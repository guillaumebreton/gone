package util

import (
	"fmt"
	"os/exec"
	"runtime"
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
	sound    string
}

// NewTimer create a new timer using a state, a command to execute,
// a painter to draw the screen and an optional sound file played on
// session end.
func NewTimer(s *state.State, p *painter.Painter, c string, n Notifier, sound string) *Timer {
	return &Timer{
		state:    s,
		painter:  p,
		command:  c,
		notifier: n,
		sound:    sound,
	}
}

// playSound plays the configured sound file in the background. It is a
// no-op when no file is set. Uses afplay on macOS and paplay elsewhere.
func playSound(file string) {
	if file == "" {
		return
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", file)
	default:
		cmd = exec.Command("paplay", file)
	}
	cmd.Start()
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
				playSound(t.sound)
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
