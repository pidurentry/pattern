package expression

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["notEqual"] = func(expression map[string]interface{}) (tools.Expression, error) {
		notEqual := &NotEqual{}
		for name, value := range expression {
			switch name {
			case "left":
				notEqual.Left = value
			case "right":
				notEqual.Right = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'notEqual' expression: %s", name))
			}
		}
		return notEqual, nil
	}
}

type NotEqual struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

func (expression *NotEqual) Test(variables tools.Variables) bool {
	return tools.LoadValue(variables, expression.Left) != tools.LoadValue(variables, expression.Right)
}
