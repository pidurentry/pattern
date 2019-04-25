package expression

import (
	"fmt"
	"github.com/pidurentry/pattern/tools"
)

func init() {
	tools.ExpressionFactory["or"] = func(expression map[string]interface{}) (tools.Expression, error) {
		_or := &Or{}
		for name, value := range expression {
			switch name {
			case "expressions":
				rawExpressions, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("'or' expression expects array of objects for 'expressions': %T", value)
				}

				expressions := make([]tools.Expression, len(rawExpressions))
				for index, rawExpression := range rawExpressions {
					expressionMap, ok := rawExpression.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("'or' expression expects array of objects for 'expressions': %T", value)
					}

					_expression, err := tools.NewExpression(expressionMap)
					if err != nil {
						return nil, err
					}

					expressions[index] = _expression
				}

				_or.Expressions = expressions
			default:
				return nil, fmt.Errorf("unknown key for 'or' expression: %s", name)
			}
		}
		return _or, nil
	}
}

type Or struct {
	Expressions []tools.Expression `json:"expressions"`
}

func (expression *Or) Test(variables tools.Variables) bool {
	for _, expression := range expression.Expressions {
		if expression.Test(variables) {
			return true
		}
	}
	return false
}
