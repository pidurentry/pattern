package action

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["rotate"] = func(action map[string]interface{}) (tools.Action, error) {
		rotate := &Rotate{}
		for name, value := range action {
			switch name {
			case "speed":
				rotate.Speed = value
			case "clockwise":
				rotate.Clockwise = value
			default:
				return nil, fmt.Errorf("unknown key for 'rotate' action: %s", name)
			}
		}
		return rotate, nil
	}
}

type Rotate struct {
	Speed     interface{} `json:"speed"`
	Clockwise interface{} `json:"clockwise"`
}

func (action *Rotate) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	clockwise := false
	if tools.LoadValue(variables, action.Speed) > 0 {
		clockwise = true
	}

	return device.Rotate(
		tools.LoadValue(variables, action.Speed),
		clockwise,
	)
}
