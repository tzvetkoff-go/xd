package xd_test

import (
	"fmt"
	"testing"

	"github.com/tzvetkoff-go/xd"
)

// ParseTest ...
type ParseTest struct {
	Where string
	Path  string
	Error error
	Value []interface{}
}

// ParseTests ...
var ParseTests = []ParseTest{
	{
		Where: Here(),
		Path:  "",
		Value: []interface{}{},
	},
	{
		Where: Here(),
		Path:  ".",
		Value: []interface{}{""},
	},
	{
		Where: Here(),
		Path:  "..",
		Value: []interface{}{"", ""},
	},
	{
		Where: Here(),
		Path:  "[0]",
		Value: []interface{}{0},
	},
	{
		Where: Here(),
		Path:  "[0].",
		Value: []interface{}{0, ""},
	},
	{
		Where: Here(),
		Path:  "[0]..",
		Value: []interface{}{0, "", ""},
	},
	{
		Where: Here(),
		Path:  "[0]..[0]",
		Value: []interface{}{0, "", "", 0},
	},
	{
		Where: Here(),
		Path:  "foo.bar.baz",
		Value: []interface{}{"foo", "bar", "baz"},
	},
	{
		Where: Here(),
		Path:  "foo[0].bar.baz",
		Value: []interface{}{"foo", 0, "bar", "baz"},
	},
	{
		Where: Here(),
		Path:  "foo[0].bar[1].baz",
		Value: []interface{}{"foo", 0, "bar", 1, "baz"},
	},
	{
		Where: Here(),
		Path:  "foo[0].bar[1].baz[-1]",
		Value: []interface{}{"foo", 0, "bar", 1, "baz", -1},
	},
	{
		Where: Here(),
		Path:  "[0xFF]",
		Value: []interface{}{0xFF},
	},
	{
		Where: Here(),
		Path:  "[0o666]",
		Value: []interface{}{0o666},
	},
	{
		Where: Here(),
		Path:  "\\.",
		Value: []interface{}{"."},
	},
	{
		Where: Here(),
		Path:  "\\..",
		Value: []interface{}{".", ""},
	},
	{
		Where: Here(),
		Path:  "\\..\\.",
		Value: []interface{}{".", "."},
	},
	{
		Where: Here(),
		Path:  "\\..\\..\\",
		Error: xd.ParseError("path ends with an escape character"),
	},
	{
		Where: Here(),
		Path:  "\\..\\..\\\\",
		Value: []interface{}{".", ".", "\\"},
	},
	{
		Where: Here(),
		Path:  "[ff]",
		Error: xd.ParseError("strconv.ParseInt: parsing \"ff\": invalid syntax"),
	},
}

// TestParse ...
func TestParse(t *testing.T) {
	for _, tt := range ParseTests {
		val, err := xd.Parse(tt.Path)

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

// ExampleParse ...
func ExampleParse() {
	if val, err := xd.Parse("foo[0].bar\\.baz"); err == nil {
		fmt.Printf("%#v\n", val)
	} else {
		panic(err)
	}

	// Output: []interface {}{"foo", 0, "bar.baz"}
}
