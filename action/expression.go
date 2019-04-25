package action

import (
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ActionFactory["expression"] = func(action map[string]interface{}) (tools.Action, error) {
		expression := &Expression{}
		for name, value := range action {
			switch name {
			case "expression":
				expressionMap, ok := value.(map[string]interface{})
				if !ok {
					return nil, errors.New("expression should be an objects")
				}

				_expression, err := tools.NewExpression(expressionMap)
				if err != nil {
					return nil, err
				}

				expression.Expression = _expression
			case "true":
				rawActions, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("'expression' action expects array of actions for 'true': %T", value)
				}

				actions, err := tools.NewActions(rawActions)
				if err != nil {
					return nil, err
				}

				expression.True = actions
			case "false":
				rawActions, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("'expression' action expects array of actions for 'false': %T", value)
				}

				actions, err := tools.NewActions(rawActions)
				if err != nil {
					return nil, err
				}

				expression.False = actions
			default:
				return nil, fmt.Errorf("unknown key for 'expression' action: %s", name)
			}
		}
		return expression, nil
	}
}

type Expression struct {
	Expression tools.Expression `json:"expression"`
	True       []tools.Action   `json:"true"`
	False      []tools.Action   `json:"false"`
}

func (action *Expression) Apply(player tools.Player, variables tools.Variables, device tools.Device) error {
	if action.Expression.Test(variables) {
		return player.QueueActions(action.True)
	} else {
		return player.QueueActions(action.False)
	}
}
