package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
	"time"
)

func init() {
	tools.ActionFactory["sleep"] = func(action map[string]interface{}) (tools.Action, error) {
		sleep := &Sleep{Unit: time.Millisecond}
		for name, value := range action {
			switch name {
			case "time":
				sleep.Time = value
			case "unit":
				unit, ok := value.(string)
				if !ok {
					return nil, errors.New(fmt.Sprintf("'sleep' action expects string for 'unit': %T", value))
				}

				duration, err := time.ParseDuration(fmt.Sprintf("1%s", unit))
				if err != nil {
					return nil, err
				}

				sleep.Unit = duration
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'sleep' action: %s", name))
			}
		}
		return sleep, nil
	}
}

type Sleep struct {
	Time interface{}   `json:"time"`
	Unit time.Duration `json:"unit"`
}

func (action *Sleep) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	return player.Sleep(
		time.Duration(tools.LoadValue(variables, action.Time)) * action.Unit,
	)
}
