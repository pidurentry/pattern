package tools

type Action interface {
	Apply(Player, Variables, Device) error
}
