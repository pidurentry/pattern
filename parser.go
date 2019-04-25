package pattern

import (
	"encoding/json"
	"fmt"
	"github.com/pidurentry/pattern/tools"
	"io"

	_ "github.com/pidurentry/pattern/action"
	_ "github.com/pidurentry/pattern/expression"
)

type Parser struct{}

func (parser *Parser) Load(data io.Reader) (Pattern, error) {
	pattern := Pattern{
		Variables: tools.NewVariables(map[string]uint64{}),
		Patterns:  make(map[string][]tools.Action),
	}
	return pattern, parser.parse(&pattern, json.NewDecoder(data))
}

func (parser *Parser) parse(pattern *Pattern, decoder *json.Decoder) error {
	var rawPattern map[string]interface{}
	if err := decoder.Decode(&rawPattern); err != nil {
		return &ParseError{fmt.Sprintf("%s", err)}
	}

	for name, value := range rawPattern {
		switch name {
		case "devices":
			if err := parser.parseDevices(pattern, value); err != nil {
				return &ParseError{fmt.Sprintf("%s", err)}
			}
		case "variables":
			if err := parser.parseVariables(pattern, value); err != nil {
				return &ParseError{fmt.Sprintf("%s", err)}
			}
		case "pattern":
			if err := parser.parsePattern(pattern, value); err != nil {
				return &ParseError{fmt.Sprintf("%s", err)}
			}
		case "patterns":
			if err := parser.parsePatterns(pattern, value); err != nil {
				return &ParseError{fmt.Sprintf("%s", err)}
			}
		default:
			return &ParseError{fmt.Sprintf("unknown modifier: %s", name)}
		}
	}

	return nil
}

func (parser *Parser) parseDevices(pattern *Pattern, rawDevices interface{}) error {
	devices, ok := rawDevices.([]interface{})
	if !ok {
		return &ParseError{fmt.Sprintf("devices should be an array of strings: %T", rawDevices)}
	}

	for _, rawDevice := range devices {
		device, ok := rawDevice.(string)
		if !ok {
			return &ParseError{fmt.Sprintf("device should be a string: %T", rawDevice)}
		}
		pattern.Devices = append(pattern.Devices, device)
	}

	return nil
}

func (parser *Parser) parseVariables(pattern *Pattern, rawVariables interface{}) error {
	variables, ok := rawVariables.(map[string]interface{})
	if !ok {
		return &ParseError{fmt.Sprintf("variables should be an map of integers: %T", rawVariables)}
	}

	for name, value := range variables {
		pattern.Variables.Initialise(name, tools.Value(value))
	}

	return nil
}

func (parser *Parser) parsePattern(pattern *Pattern, rawName interface{}) error {
	name, ok := rawName.(string)
	if !ok {
		return &ParseError{fmt.Sprintf("pattern should be a string: %T", rawName)}
	}

	pattern.Pattern = name
	return nil
}

func (parser *Parser) parsePatterns(pattern *Pattern, rawPatterns interface{}) error {
	patterns, ok := rawPatterns.(map[string]interface{})
	if !ok {
		return &ParseError{fmt.Sprintf("patterns should be a map of arrays: %T", rawPatterns)}
	}

	for name, rawActions := range patterns {
		actions, err := parser.parseActions(rawActions)
		if err != nil {
			return err
		}
		pattern.Patterns[name] = actions
	}

	return nil
}

func (parser *Parser) parseActions(rawActions interface{}) ([]tools.Action, error) {
	actions, ok := rawActions.([]interface{})
	if !ok {
		return nil, &ParseError{fmt.Sprintf("actions should be an array maps: %T", rawActions)}
	}
	return tools.NewActions(actions)
}
