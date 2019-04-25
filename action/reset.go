package action

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["reset"] = func(action map[string]interface{}) (tools.Action, error) {
		reset := &Reset{}
		for name, value := range action {
			switch name {
			case "variable":
				variable, ok := value.(string)
				if !ok {
					return nil, fmt.Errorf("'reset' action expects string for 'variable': %T", value)
				}
				reset.Variable = variable
			default:
				return nil, fmt.Errorf("unknown key for 'reset' action: %s", name)
			}
		}
		return reset, nil
	}
}

type Reset struct {
	Variable string `json:"variable"`
}

func (action *Reset) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	variables.Reset(action.Variable)
	return nil
}
