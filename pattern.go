package pattern

type Pattern struct {
	Devices   []string                 `json:"devices"`
	Variables Variables                `json:"variables"`
	Pattern   string                   `json:"pattern"`
	Patterns  map[string][]interface{} `json:"patterns"`
}
