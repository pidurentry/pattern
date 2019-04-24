package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["rotate"] = func(action map[string]interface{}) (interface{}, error) {
		rotate := &Rotate{}
		for name, value := range action {
			switch name {
			case "speed":
				rotate.Speed = value
			case "clockwise":
				rotate.Clockwise = value
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'rotate' action: %s", name))
			}
		}
		return rotate, nil
	}
}

type Rotate struct {
	Speed     interface{} `json:"speed"`
	Clockwise interface{} `json:"clockwise"`
}
