package expression

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["and"] = func(expression map[string]interface{}) (tools.Expression, error) {
		_and := &And{}
		for name, value := range expression {
			switch name {
			case "expressions":
				rawExpressions, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("'and' expression expects array of objects for 'expressions': %T", value)
				}

				expressions := make([]tools.Expression, len(rawExpressions))
				for index, rawExpression := range rawExpressions {
					expressionMap, ok := rawExpression.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("'and' expression expects array of objects for 'expressions': %T", value)
					}

					_expression, err := tools.NewExpression(expressionMap)
					if err != nil {
						return nil, err
					}

					expressions[index] = _expression
				}

				_and.Expressions = expressions
			default:
				return nil, fmt.Errorf("unknown key for 'and' expression: %s", name)
			}
		}
		return _and, nil
	}
}

type And struct {
	Expressions []tools.Expression `json:"expressions"`
}

func (expression *And) Test(variables tools.Variables) bool {
	for _, expression := range expression.Expressions {
		if !expression.Test(variables) {
			return false
		}
	}
	return true
}
