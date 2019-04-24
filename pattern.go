package pattern

import "github.com/pidurentry/pattern/tools"

type Pattern struct {
	Devices   []string                 `json:"devices"`
	Variables tools.Variables          `json:"variables"`
	Pattern   string                   `json:"pattern"`
	Patterns  map[string][]interface{} `json:"patterns"`
}
