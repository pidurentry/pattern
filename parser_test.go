package pattern

import (
	"bufio"
	"errors"
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
		return errors.New("pattern do not match")
	}
	return nil
}

func comparePatterns(pattern, expected Pattern) error {
	// TODO: Figure out how to compare actions
	return nil
}
