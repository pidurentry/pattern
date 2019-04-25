package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["move"] = func(action map[string]interface{}) (tools.Action, error) {
		move := &Move{}
		for name, value := range action {
			switch name {
			case "value":
				move.Value = value
			case "speed":
				move.Speed = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'move' action: %s", name))
			}
		}
		return move, nil
	}
}

type Move struct {
	Value interface{} `json:"value"`
	Speed interface{} `json:"speed"`
}

func (action *Move) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	return device.Move(
		tools.LoadValue(variables, action.Value),
		tools.LoadValue(variables, action.Speed),
	)
}
