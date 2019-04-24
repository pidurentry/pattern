package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["repeat"] = func(action map[string]interface{}) (interface{}, error) {
		repeat := &Repeat{}
		for name, value := range action {
			switch name {
			case "count":
				repeat.Count = value
			case "pattern":
				rawActions, ok := value.([]interface{})
				if !ok {
					return nil, errors.New(fmt.Sprintf("'repeat' action expects array of actions for 'pattern': %T", value))
				}

				actions, err := tools.NewActions(rawActions)
				if err != nil {
					return nil, err
				}

				repeat.Pattern = actions
			default:
				return nil, errors.New(fmt.Sprintf("unknown key for 'repeat' action: %s", name))
			}
		}
		return repeat, nil
	}
}

type Repeat struct {
	Count   interface{}   `json:"count"`
	Pattern []interface{} `json:"pattern"`
}
