package tools

type Expression interface {
	Test(Variables) bool
}
