package expression

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["equal"] = func(expression map[string]interface{}) (tools.Expression, error) {
		equal := &Equal{}
		for name, value := range expression {
			switch name {
			case "left":
				equal.Left = value
			case "right":
				equal.Right = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'equal' expression: %s", name))
			}
		}
		return equal, nil
	}
}

type Equal struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *Equal) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) == tools.LoadValue(variables, expression.Right)
}
