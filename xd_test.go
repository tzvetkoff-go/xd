package xd_test

import (
	"fmt"
	"testing"

	"github.com/tzvetkoff-go/xd"
)

// DigTest ...
type DigTest struct {
	Where string
	Root  interface{}
	Path  string
	Error error
	Value interface{}
}

// DigTests ...
var DigTests = []DigTest{
	{
		Where: Here(),
		Root:  map[string]interface{}{},
		Path:  "[0]",
		Error: xd.NotArrayError(""),
	},
	{
		Where: Here(),
		Root:  []interface{}{},
		Path:  ".",
		Error: xd.NotMapError(""),
	},
	{
		Where: Here(),
		Root:  map[string]interface{}{},
		Path:  ".",
		Error: xd.NotFoundError(""),
	},
	{
		Where: Here(),
		Root: map[string]interface{}{
			"key": "value",
		},
		Path:  "key",
		Value: "value",
	},
	{
		Where: Here(),
		Root:  []interface{}{0, 1, 2},
		Path:  "[0]",
		Value: 0,
	},
	{
		Where: Here(),
		Root:  []interface{}{0, 1, 2},
		Path:  "[-1]",
		Value: 2,
	},
	{
		Where: Here(),
		Root: map[string]interface{}{
			"arr": []interface{}{
				"object",
				map[string]interface{}{
					"type": "something",
				},
			},
		},
		Path:  "arr[1].type",
		Value: "something",
	},
}

// TestDig ...
func TestDigE(t *testing.T) {
	for _, tt := range DigTests {
		val, err := xd.DigE(tt.Root, tt.Path)

		if tt.Error != nil {
			if !Compare(err, tt.Error) {
				t.Errorf("\n  %s:\n        expected error: %#v,\n             got error: %#v", tt.Where, tt.Error, err)
			}
		} else {
			if !Compare(val, tt.Value) {
				t.Errorf("\n  %s:\n        expected value: %#v,\n             got value: %#v", tt.Where, tt.Value, val)
			}
		}
	}
}

// ExampleDig ...
func ExampleDigE() {
	m := map[string]interface{}{
		"arr": []interface{}{
			"object",
			map[string]interface{}{
				"foo.bar": "baz",
			},
		},
	}

	if val, err := xd.DigE(m, "arr[-1].foo\\.bar"); err == nil {
		fmt.Printf("%#v\n", val)
	} else {
		panic(err)
	}

	// Output: "baz"
}
