package pattern

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/pidurentry/pattern/action"
	"github.com/pidurentry/pattern/expression"
	"github.com/pidurentry/pattern/tools"
	"os"
	"sync/atomic"
	"testing"
)

func TestParseExpansion(t *testing.T) {
	testParse(
		t,
		"examples/expansion.json",
		Pattern{
			Devices: []string{"FleshlightLaunch", "Linear"},
			Variables: tools.NewVariables(map[string]uint64{
				"counter": 0,
			}),
			Pattern: "level1",
			Patterns: map[string][]tools.Action{
				"level1": []tools.Action{
					&action.Repeat{Count: 5, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 10},
						&action.Move{Value: 99, Speed: 10},
					}},
					&action.Repeat{Count: 1, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 20},
						&action.Move{Value: 99, Speed: 20},
					}},
					&action.Expression{
						Expression: &expression.GreaterThanOrEqual{Left: "counter", Right: 10},
						True: []tools.Action{
							&action.Reset{Variable: "counter"},
							&action.Goto{Pattern: "level2"},
						},
						False: []tools.Action{
							&action.Increment{Variable: "counter", Value: 1},
						},
					},
				},
				"level2": []tools.Action{
					&action.Repeat{Count: 10, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 10},
						&action.Move{Value: 99, Speed: 10},
					}},
					&action.Repeat{Count: 2, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 20},
						&action.Move{Value: 99, Speed: 20},
					}},
					&action.Repeat{Count: 1, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 30},
						&action.Move{Value: 99, Speed: 30},
					}},
					&action.Expression{
						Expression: &expression.GreaterThanOrEqual{Left: "counter", Right: 10},
						True: []tools.Action{
							&action.Reset{Variable: "counter"},
							&action.Goto{Pattern: "level3"},
						},
						False: []tools.Action{
							&action.Increment{Variable: "counter", Value: 1},
						},
					},
				},
				"level3": []tools.Action{
					&action.Repeat{Count: 15, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 10},
						&action.Move{Value: 99, Speed: 10},
					}},
					&action.Repeat{Count: 5, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 20},
						&action.Move{Value: 99, Speed: 20},
					}},
					&action.Repeat{Count: 2, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 30},
						&action.Move{Value: 99, Speed: 30},
					}},
					&action.Repeat{Count: 1, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 40},
						&action.Move{Value: 99, Speed: 40},
					}},
					&action.Expression{
						Expression: &expression.GreaterThanOrEqual{Left: "counter", Right: 10},
						True: []tools.Action{
							&action.Reset{Variable: "counter"},
							&action.Goto{Pattern: "level4"},
						},
						False: []tools.Action{
							&action.Increment{Variable: "counter", Value: 1},
						},
					},
				},
				"level4": []tools.Action{
					&action.Repeat{Count: 20, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 10},
						&action.Move{Value: 99, Speed: 10},
					}},
					&action.Repeat{Count: 10, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 20},
						&action.Move{Value: 99, Speed: 20},
					}},
					&action.Repeat{Count: 5, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 30},
						&action.Move{Value: 99, Speed: 30},
					}},
					&action.Repeat{Count: 2, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 40},
						&action.Move{Value: 99, Speed: 40},
					}},
					&action.Repeat{Count: 1, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 50},
						&action.Move{Value: 99, Speed: 50},
					}},
					&action.Expression{
						Expression: &expression.GreaterThanOrEqual{Left: "counter", Right: 10},
						True: []tools.Action{
							&action.Reset{Variable: "counter"},
							&action.Goto{Pattern: "level5"},
						},
						False: []tools.Action{
							&action.Increment{Variable: "counter", Value: 1},
						},
					},
				},
				"level5": []tools.Action{
					&action.Repeat{Count: 25, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 10},
						&action.Move{Value: 99, Speed: 10},
					}},
					&action.Repeat{Count: 15, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 20},
						&action.Move{Value: 99, Speed: 20},
					}},
					&action.Repeat{Count: 10, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 30},
						&action.Move{Value: 99, Speed: 30},
					}},
					&action.Repeat{Count: 5, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 40},
						&action.Move{Value: 99, Speed: 40},
					}},
					&action.Repeat{Count: 2, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 50},
						&action.Move{Value: 99, Speed: 50},
					}},
					&action.Repeat{Count: 1, Pattern: []tools.Action{
						&action.Move{Value: 0, Speed: 60},
						&action.Move{Value: 99, Speed: 60},
					}},
				},
			},
		},
	)
}

func testParse(t *testing.T, filename string, expected Pattern) {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	pattern, err := (&Parser{}).Load(bufio.NewReader(file))
	if err != nil {
		t.Fatal(err)
	}

	if err := comparePatternDevices(pattern, expected); err != nil {
		t.Fatal(err)
	}

	if err := compareVariables(pattern, expected); err != nil {
		t.Fatal(err)
	}

	if err := comparePattern(pattern, expected); err != nil {
		t.Fatal(err)
	}

	if err := comparePatterns(pattern, expected); err != nil {
		t.Fatal(err)
	}
}

func comparePatternDevices(pattern, expected Pattern) error {
	if len(pattern.Devices) != len(expected.Devices) {
		return errors.New("pattern devices do not match")
	}

	devices := make(map[string]bool)
	for _, device := range pattern.Devices {
		devices[device] = true
	}

	for _, device := range expected.Devices {
		_, ok := devices[device]
		if !ok {
			return errors.New("pattern devices do not match")
		}
		delete(devices, device)
	}

	if len(devices) > 0 {
		return errors.New("pattern devices do not match")
	}

	return nil
}

func compareVariables(pattern, expected Pattern) error {
	variables := pattern.Variables.FetchAll()

	for variable, expectedValue := range expected.Variables.FetchAll() {
		value, ok := variables[variable]
		if !ok {
			return errors.New("pattern variables do not match")
		}
		delete(variables, variable)

		if atomic.LoadUint64(value) != atomic.LoadUint64(expectedValue) {
			return errors.New("pattern variables do not match")
		}
	}

	if len(variables) > 0 {
		return errors.New("pattern variables do not match")
	}

	return nil
}

func comparePattern(pattern, expected Pattern) error {
	if pattern.Pattern != expected.Pattern {
		return errors.New("pattern does not match")
	}
	return nil
}

func comparePatterns(pattern, expected Pattern) error {
	if len(pattern.Patterns) != len(expected.Patterns) {
		return errors.New("pattern patterns does not match")
	}

	patterns := pattern.Patterns
	for name, expectedActions := range expected.Patterns {
		actions, ok := patterns[name]
		if !ok {
			return errors.New("pattern patterns does not match")
		}
		delete(patterns, name)

		if err := compareActions(actions, expectedActions); err != nil {
			return err
		}
	}

	if len(patterns) > 0 {
		return errors.New("pattern patterns does not match")
	}

	return nil
}

func compareActions(actions, expectedActions []tools.Action) error {
	if len(actions) != len(expectedActions) {
		return errors.New("actions do not match")
	}

	for index, _action := range actions {
		expectedAction := expectedActions[index]

		if fmt.Sprintf("%T", _action) != fmt.Sprintf("%T", expectedAction) {
			return errors.New("action does not match")
		}

		switch _action := _action.(type) {
		case *action.Decrement:
			if err := compareDecrementActions(_action, expectedAction.(*action.Decrement)); err != nil {
				return err
			}
		case *action.Expression:
			if err := compareExpressionActions(_action, expectedAction.(*action.Expression)); err != nil {
				return err
			}
		case *action.Goto:
			if err := compareGotoActions(_action, expectedAction.(*action.Goto)); err != nil {
				return err
			}
		case *action.Increment:
			if err := compareIncrementActions(_action, expectedAction.(*action.Increment)); err != nil {
				return err
			}
		case *action.Move:
			if err := compareMoveActions(_action, expectedAction.(*action.Move)); err != nil {
				return err
			}
		case *action.Repeat:
			if err := compareRepeatActions(_action, expectedAction.(*action.Repeat)); err != nil {
				return err
			}
		case *action.Reset:
			if err := compareResetActions(_action, expectedAction.(*action.Reset)); err != nil {
				return err
			}
		default:
			return fmt.Errorf("no comparison for '%T' action", _action)
		}
	}

	return nil
}

func compareDecrementActions(decrement, expected *action.Decrement) error {
	if decrement.Variable != expected.Variable {
		return errors.New("decrement action variable does not match")
	}

	switch value := decrement.Value.(type) {
	case string:
		if value != expected.Value.(string) {
			return errors.New("decrement action value does not match")
		}
	default:
		if tools.Value(value) != tools.Value(expected.Value) {
			return errors.New("decrement action value does not match")
		}
	}

	return nil
}

func compareExpressionActions(expression, expected *action.Expression) error {
	if err := compareExpression(expression.Expression, expected.Expression); err != nil {
		return err
	}

	if err := compareActions(expression.True, expected.True); err != nil {
		return err
	}

	if err := compareActions(expression.False, expected.False); err != nil {
		return err
	}

	return nil
}

func compareExpression(_expression, expected tools.Expression) error {
	if fmt.Sprintf("%T", _expression) != fmt.Sprintf("%T", expected) {
		return errors.New("expression does not match")
	}

	switch _expression := _expression.(type) {
	case *expression.GreaterThanOrEqual:
		if err := compareGreaterThanOrEqualExpression(_expression, expected.(*expression.GreaterThanOrEqual)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("no comparison for '%T' expression", _expression)
	}

	return nil
}

func compareGreaterThanOrEqualExpression(greaterThanOrEqual, expected *expression.GreaterThanOrEqual) error {
	switch left := greaterThanOrEqual.Left.(type) {
	case string:
		if left != expected.Left.(string) {
			return errors.New("greaterThanOrEqual expression left does not match")
		}
	default:
		if tools.Value(left) != tools.Value(expected.Left) {
			return errors.New("greaterThanOrEqual expression left does not match")
		}
	}

	switch right := greaterThanOrEqual.Right.(type) {
	case string:
		if right != expected.Right.(string) {
			return errors.New("greaterThanOrEqual expression right does not match")
		}
	default:
		if tools.Value(right) != tools.Value(expected.Right) {
			return errors.New("greaterThanOrEqual expression right does not match")
		}
	}

	return nil
}

func compareGotoActions(_goto, expected *action.Goto) error {
	if _goto.Pattern != expected.Pattern {
		return errors.New("goto action pattern does not match")
	}

	return nil
}

func compareIncrementActions(increment, expected *action.Increment) error {
	if increment.Variable != expected.Variable {
		return errors.New("decrement action variable does not match")
	}

	switch value := increment.Value.(type) {
	case string:
		if value != expected.Value.(string) {
			return errors.New("increment action value does not match")
		}
	default:
		if tools.Value(value) != tools.Value(expected.Value) {
			return errors.New("increment action value does not match")
		}
	}

	return nil
}

func compareMoveActions(move, expected *action.Move) error {
	switch value := move.Value.(type) {
	case string:
		if value != expected.Value.(string) {
			return errors.New("move action value does not match")
		}
	default:
		if tools.Value(value) != tools.Value(expected.Value) {
			return errors.New("move action value does not match")
		}
	}

	switch speed := move.Speed.(type) {
	case string:
		if speed != expected.Speed.(string) {
			return errors.New("move action speed does not match")
		}
	default:
		if tools.Value(speed) != tools.Value(expected.Speed) {
			return errors.New("move action speed does not match")
		}
	}

	return nil
}

func compareRepeatActions(repeat, expected *action.Repeat) error {
	switch count := repeat.Count.(type) {
	case string:
		if count != expected.Count.(string) {
			return errors.New("repeat action count does not match")
		}
	default:
		if tools.Value(count) != tools.Value(expected.Count) {
			return errors.New("repeat action count does not match")
		}
	}

	if err := compareActions(repeat.Pattern, expected.Pattern); err != nil {
		return err
	}

	return nil
}

func compareResetActions(reset, expected *action.Reset) error {
	if reset.Variable != expected.Variable {
		return errors.New("reset action variable does not match")
	}

	return nil
}
