package expression

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["greaterThanOrEqual"] = func(expression map[string]interface{}) (tools.Expression, error) {
		greaterThanOrEqual := &GreaterThanOrEqual{}
		for name, value := range expression {
			switch name {
			case "left":
				greaterThanOrEqual.Left = value
			case "right":
				greaterThanOrEqual.Right = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'greaterThanOrEqual' expression: %s", name))
			}
		}
		return greaterThanOrEqual, nil
	}
}

type GreaterThanOrEqual struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *GreaterThanOrEqual) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) >= tools.LoadValue(variables, expression.Right)
}
