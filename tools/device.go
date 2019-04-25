package tools

type Device interface {
	Move(Value, Speed uint64) error
	Rotate(Speed uint64, Clockwise bool) error
	Vibrate(Speed uint64) error
}
