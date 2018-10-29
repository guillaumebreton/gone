package util

import (
	"github.com/0xAX/notificator"
)

// Notifier is an interface to send notifications.
type Notifier interface {
	Notify(title string, text string) error
}

// desktopNotifier implements the Notifier interface.
type desktopNotifier struct {
	*notificator.Notificator
}

// NewDesktopNotifier creates a new notifier.
func NewDesktopNotifier() Notifier {
	return &desktopNotifier{
		notificator.New(notificator.Options{
			DefaultIcon: "icon/default.png",
			AppName:     "gone (pomodoro)",
		}),
	}
}

// Notify sends a desktop notification.
func (d *desktopNotifier) Notify(title, text string) error {
	return d.Push(title, text, "", notificator.UR_NORMAL)
}

// nullNotifier is a notifier that doesn't do anything.
type nullNotifier struct{}

// NewNullNotifier returns a notifier that is silent.
func NewNullNotifier() Notifier {
	return &nullNotifier{}
}

// Notify does not do anything.
func (n *nullNotifier) Notify(title, text string) error {
	return nil
}
