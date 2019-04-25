package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
	"sync/atomic"
)

func init() {
	tools.ActionFactory["increment"] = func(action map[string]interface{}) (tools.Action, error) {
		increment := &Increment{Value: 1}
		for name, value := range action {
			switch name {
			case "variable":
				variable, ok := value.(string)
				if !ok {
					return nil, errors.New(fmt.Sprintf("'increment' action expects string for 'variable': %T", value))
				}
				increment.Variable = variable
			case "value":
				increment.Value = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'increment' action: %s", name))
			}
		}
		return increment, nil
	}
}

type Increment struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
}

func (action *Increment) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	variable := variables.Fetch(action.Variable)
	atomic.AddUint64(variable, tools.LoadValue(variables, action.Value))
	return nil
}
