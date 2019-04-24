package action

type Increment struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
}
