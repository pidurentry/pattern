package pattern

import "io"

type Loader interface {
	Load(io.Reader) (Pattern, error)
}
