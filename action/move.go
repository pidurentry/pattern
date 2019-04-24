package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["move"] = func(action map[string]interface{}) (interface{}, error) {
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
