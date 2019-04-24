package action

type If struct {
	Expression interface{}   `json:"expression"`
	If         []interface{} `json:"if"`
	Else       []interface{} `json:"else"`
}
