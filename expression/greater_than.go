package expression

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["greaterThan"] = func(expression map[string]interface{}) (tools.Expression, error) {
		greaterThan := &GreaterThan{}
		for name, value := range expression {
			switch name {
			case "left":
				greaterThan.Left = value
			case "right":
				greaterThan.Right = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'greaterThan' expression: %s", name))
			}
		}
		return greaterThan, nil
	}
}

type GreaterThan struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *GreaterThan) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) > tools.LoadValue(variables, expression.Right)
}
