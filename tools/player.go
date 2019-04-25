package tools

import "time"

type Player interface {
	Goto(string) error
	QueueActions([]Action) error
	Sleep(time.Duration) error
}
