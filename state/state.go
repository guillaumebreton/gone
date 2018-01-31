package state

import "fmt"

type Callback func()

const (
	RUNNING = iota
	PAUSED  = iota
	WAITING = iota
)

type State struct {
	currentState          int
	currentIdx            int
	duration              int
	durations             map[byte]int
	pattern               string
	paused                bool
	confirm               bool
	resumeWaitingCallback Callback
}

func NewState(pattern string, w, s, l int) *State {
	durations := map[byte]int{'w': w * 60, 's': s * 60, 'l': l * 60}
	start := pattern[0]
	return &State{
		currentState: RUNNING,
		currentIdx:   0,
		duration:     durations[start],
		durations:    durations,
		pattern:      pattern,
		paused:       false,
		confirm:      false,
	}
}

func (s *State) Pause() {
	s.currentState = PAUSED
}

func (s *State) Resume() {
	if s.currentState == WAITING {
		//call the call back
		s.Next()
		go s.resumeWaitingCallback()
		s.resumeWaitingCallback = nil
	}
	s.currentState = RUNNING
}

func (s *State) WaitForConfirm(callback Callback) {
	s.currentState = WAITING
	s.resumeWaitingCallback = callback
}

func (s *State) IsRunning() bool {
	return s.currentState == RUNNING
}

func (s *State) IsWaiting() bool {
	return s.currentState == WAITING
}

func (s *State) IsEnded() bool {
	return s.duration == 0
}

func (s *State) Decrease() {
	if s.duration > 0 {
		s.duration--
	}
}

func (s *State) Next() {
	if s.currentIdx == len(s.pattern)-1 {
		s.currentIdx = 0
	} else {
		s.currentIdx++
	}
	s.duration = s.durations[s.pattern[s.currentIdx]]
}

func (s *State) Message() string {
	if s.IsRunning() {
		switch s.pattern[s.currentIdx] {
		case 'w':
			return "working session"
		case 's':
			return "short break"
		case 'l':
			return "long break"
		case 'c':
			return "Continue? (y/n)"
		}
	} else if s.currentState == WAITING {
		return "continue? [y/n]"
	}
	return "paused"
}

// StatusMessage returns a message describing the state of a session.
func (s *State) StatusMessage() string {
	msgs := map[byte]string{
		'w': "working session",
		's': "short break",
		'l': "long break",
	}

	if !s.IsEnded() {
		return fmt.Sprintf("A %s is in progress", msgs[s.pattern[s.currentIdx]])
	}

	return fmt.Sprintf("Your %s has ended, time for a %s.", msgs[s.pattern[s.currentIdx]], msgs[s.pattern[s.currentIdx+1]])
}

// Duration format a duration to 12:34 (mm:ss).
func (s *State) Duration() string {
	seconds := s.duration % 60
	minutes := s.duration / 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
