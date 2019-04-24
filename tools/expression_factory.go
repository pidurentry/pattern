package tools

import (
	"errors"
	"fmt"
)

type expressionFactory func(map[string]interface{}) (Expression, error)

var ExpressionFactory = make(map[string]expressionFactory)

func NewExpression(expressionMap map[string]interface{}) (Expression, error) {
	expressionType, ok := expressionMap["type"].(string)
	if !ok {
		return nil, errors.New("expression type should be a string")
	}
	delete(expressionMap, "type")

	factory, ok := ExpressionFactory[expressionType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown expression: %s", expressionType))
	}

	return factory(expressionMap)
}
