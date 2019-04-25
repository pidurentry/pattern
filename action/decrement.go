package action

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
	"sync/atomic"
)

func init() {
	tools.ActionFactory["decrement"] = func(action map[string]interface{}) (tools.Action, error) {
		decrement := &Decrement{Value: 1}
		for name, value := range action {
			switch name {
			case "variable":
				variable, ok := value.(string)
				if !ok {
					return nil, fmt.Errorf("'decrement' action expects string for 'variable': %T", value)
				}
				decrement.Variable = variable
			case "value":
				decrement.Value = value
			default:
				return nil, fmt.Errorf("unknown key for 'decrement' action: %s", name)
			}
		}
		return decrement, nil
	}
}

type Decrement struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
}

func (action *Decrement) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	variable := variables.Fetch(action.Variable)
	atomic.AddUint64(variable, 0-tools.LoadValue(variables, action.Value))
	return nil
}
