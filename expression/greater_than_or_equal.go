package expression

import "github.com/pidurentry/pattern"

type GreaterThanOrEqual struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *GreaterThanOrEqual) Test(variables pattern.Variables) bool {
	return pattern.LoadValue(variables, expression.Left) >= pattern.LoadValue(variables, expression.Right)
}
