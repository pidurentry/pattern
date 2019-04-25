package action

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["vibrate"] = func(action map[string]interface{}) (tools.Action, error) {
		vibrate := &Vibrate{}
		for name, value := range action {
			switch name {
			case "speed":
				vibrate.Speed = value
			default:
				return nil, fmt.Errorf("unknown key for 'vibrate' action: %s", name)
			}
		}
		return vibrate, nil
	}
}

type Vibrate struct {
	Speed interface{} `json:"speed"`
}

func (action *Vibrate) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	return device.Vibrate(
		tools.LoadValue(variables, action.Speed),
	)
}
