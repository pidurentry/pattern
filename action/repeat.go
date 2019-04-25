package action

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["repeat"] = func(action map[string]interface{}) (tools.Action, error) {
		repeat := &Repeat{}
		for name, value := range action {
			switch name {
			case "count":
				repeat.Count = value
			case "pattern":
				rawActions, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("'repeat' action expects array of actions for 'pattern': %T", value)
				}

				actions, err := tools.NewActions(rawActions)
				if err != nil {
					return nil, err
				}

				repeat.Pattern = actions
			default:
				return nil, fmt.Errorf("unknown key for 'repeat' action: %s", name)
			}
		}
		return repeat, nil
	}
}

type Repeat struct {
	Count   interface{}    `json:"count"`
	Pattern []tools.Action `json:"pattern"`
}

func (action *Repeat) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	count := tools.LoadValue(variables, action.Count)
	for i := uint64(0); i < count; i++ {
		if err := player.QueueActions(action.Pattern); err != nil {
			return err
		}
	}
	return nil
}
