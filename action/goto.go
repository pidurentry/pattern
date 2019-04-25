package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["goto"] = func(action map[string]interface{}) (tools.Action, error) {
		_goto := &Goto{}
		for name, value := range action {
			switch name {
			case "pattern":
				pattern, ok := value.(string)
				if !ok {
					return nil, errors.New(fmt.Sprintf("'goto' action expects string for 'pattern': %T", value))
				}
				_goto.Pattern = pattern
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'goto' action: %s", name))
			}
		}
		return _goto, nil
	}
}

type Goto struct {
	Pattern string `json:"pattern"`
}

func (action *Goto) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	return player.Goto(action.Pattern)
}
