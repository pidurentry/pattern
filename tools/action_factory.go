package tools

import (
	"errors"
	"fmt"
)

type actionFactory func(map[string]interface{}) (Action, error)

var ActionFactory = make(map[string]actionFactory)

func NewActions(rawActions []interface{}) ([]Action, error) {
	actions := make([]Action, len(rawActions))
	for index, rawAction := range rawActions {
		actionMap, ok := rawAction.(map[string]interface{})
		if !ok {
			return nil, errors.New("actions should be a map of objects")
		}

		action, err := NewAction(actionMap)
		if err != nil {
			return nil, err
		}

		actions[index] = action
	}
	return actions, nil
}

func NewAction(actionMap map[string]interface{}) (Action, error) {
	actionType, ok := actionMap["action"].(string)
	if !ok {
		return nil, errors.New("action type should be a string")
	}
	delete(actionMap, "action")

	factory, ok := ActionFactory[actionType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown action: %s", actionType))
	}

	return factory(actionMap)
}
