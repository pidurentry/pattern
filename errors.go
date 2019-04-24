package pattern

import "fmt"

type ParseError struct {
	Type string
}

func (error *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", error.Type)
}
