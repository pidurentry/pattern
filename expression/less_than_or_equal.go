package expression

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["lessThanOrEqual"] = func(expression map[string]interface{}) (tools.Expression, error) {
		lessThanOrEqual := &LessThanOrEqual{}
		for name, value := range expression {
			switch name {
			case "left":
				lessThanOrEqual.Left = value
			case "right":
				lessThanOrEqual.Right = value
			default:
				return nil, fmt.Errorf("unknown key for 'lessThanOrEqual' expression: %s", name)
			}
		}
		return lessThanOrEqual, nil
	}
}

type LessThanOrEqual struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *LessThanOrEqual) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) <= tools.LoadValue(variables, expression.Right)
}
