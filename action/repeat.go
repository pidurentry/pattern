package action

type Repeat struct {
	Count   interface{}   `json:"count"`
	Pattern []interface{} `json:"pattern"`
}
