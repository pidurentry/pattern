package pattern

type Variables interface {
	Fetch(string) *uint64
}

type variables map[string]*uint64

func (variables variables) Fetch(variable string) *uint64 {
	value, ok := variables[variable]
	if !ok {
		new := uint64(0)
		value = &new
		variables[variable] = value
	}
	return value
}
