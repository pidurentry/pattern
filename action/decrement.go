package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["decrement"] = func(action map[string]interface{}) (interface{}, error) {
		decrement := &Decrement{Value: 1}
		for name, value := range action {
			switch name {
			case "variable":
				variable, ok := value.(string)
				if !ok {
					return nil, errors.New(fmt.Sprintf("'decrement' action expects string for 'variable': %T", value))
				}
				decrement.Variable = variable
			case "value":
				decrement.Value = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'decrement' action: %s", name))
			}
		}
		return decrement, nil
	}
}

type Decrement struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
}
