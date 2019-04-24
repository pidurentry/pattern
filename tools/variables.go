package tools

import "sync/atomic"

type Variables interface {
	Initialise(string, uint64)
	Fetch(string) *uint64
	FetchAll() map[string]*uint64
	Reset(string) *uint64
}

type variables struct {
	current  map[string]*uint64
	original map[string]uint64
}

func NewVariables(values map[string]uint64) Variables {
	variables := &variables{
		current:  make(map[string]*uint64),
		original: make(map[string]uint64),
	}

	for variable, value := range values {
		variables.Initialise(variable, value)
	}

	return variables
}

func (variables *variables) Initialise(variable string, value uint64) {
	variables.current[variable] = &value
	variables.original[variable] = value
}

func (variables *variables) Fetch(variable string) *uint64 {
	value, ok := variables.current[variable]
	if !ok {
		return variables.Reset(variable)
	}
	return value
}

func (variables *variables) FetchAll() map[string]*uint64 {
	return variables.current
}

func (variables *variables) Reset(variable string) *uint64 {
	value, ok := variables.original[variable]
	if !ok {
		value = 0
	}
	atomic.StoreUint64(variables.current[variable], value)
	return variables.current[variable]
}
