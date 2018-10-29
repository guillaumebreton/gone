package state

import "testing"

func TestStatusMessage(t *testing.T) {
	tests := []struct {
		name string
		got  *State
		want string
	}{
		{
			name: "default from main",
			got:  NewState("wswswl", 25, 5, 15),
			want: "A working session is in progress",
		},
		{
			name: "one working state",
			got:  NewState("w", 25, 5, 15),
			want: "A working session is in progress",
		},
		{
			name: "one break state",
			got:  NewState("s", 25, 5, 15),
			want: "A short break is in progress",
		},
		{
			name: "next state",
			got: func() *State {
				s := NewState("wswl", 25, 5, 15)
				s.currentIdx++
				return s
			}(),
			want: "A short break is in progress",
		},
		{
			name: "invalid state",
			got: func() *State {
				s := NewState("s", 25, 5, 15)
				s.currentIdx++
				return s
			}(),
			want: "Gone encountered an unknown state",
		},
		{
			name: "ended state",
			got: func() *State {
				s := NewState("wswl", 25, 5, 15)
				s.duration = 0
				return s
			}(),
			want: "Your working session has ended, time for a short break.",
		},
		{
			name: "ended state second index",
			got: func() *State {
				s := NewState("wswl", 25, 5, 15)
				s.currentIdx++
				s.duration = 0
				return s
			}(),
			want: "Your short break has ended, time for a working session.",
		},
		{
			name: "ended state circular state",
			got: func() *State {
				s := NewState("wswl", 25, 5, 15)
				s.currentIdx = 3
				s.duration = 0
				return s
			}(),
			want: "Your long break has ended, time for a working session.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			got := test.got.StatusMessage()
			if got != test.want {
				t.Errorf("want status = %q, got %q", test.want, got)
			}
		})
	}
}
