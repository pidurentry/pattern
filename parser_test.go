package pattern

import (
	"bufio"
	"fmt"
	"github.com/pidurentry/pattern/tools"
	"os"
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

	// TODO: Check pattern equals the expected pattern
	fmt.Printf("%#v\n", pattern)
}
