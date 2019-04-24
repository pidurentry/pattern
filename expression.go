package pattern

type Expression interface {
	Test(Variables) bool
}
