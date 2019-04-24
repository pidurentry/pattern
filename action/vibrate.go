package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["vibrate"] = func(action map[string]interface{}) (interface{}, error) {
		vibrate := &Vibrate{}
		for name, value := range action {
			switch name {
			case "speed":
				vibrate.Speed = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'vibrate' action: %s", name))
			}
		}
		return vibrate, nil
	}
}

type Vibrate struct {
	Speed interface{} `json:"speed"`
}
