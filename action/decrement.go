package action

type Decrement struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
}
