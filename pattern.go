package pattern

type Pattern struct {
	Devices   []string                 `json:"devices"`
	Variables map[string]uint64        `json:"variables"`
	Pattern   string                   `json:"pattern"`
	Patterns  map[string][]interface{} `json:"patterns"`
}
