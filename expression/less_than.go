package expression

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["lessThan"] = func(expression map[string]interface{}) (tools.Expression, error) {
		lessThan := &LessThan{}
		for name, value := range expression {
			switch name {
			case "left":
				lessThan.Left = value
			case "right":
				lessThan.Right = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'lessThan' expression: %s", name))
			}
		}
		return lessThan, nil
	}
}

type LessThan struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *LessThan) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) < tools.LoadValue(variables, expression.Right)
}
