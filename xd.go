package xd

import (
	"fmt"
)

//
// Dig returns the value at a given path or nil.
// All sorts of errors will also result in nil.
//
func Dig(root interface{}, path string) interface{} {
	val, _ := DigE(root, path)
	return val
}

//
// DigB returns the value at a given path or nil as first result, and true or false as second.
// All sorts of errors will also result in nil, false.
//
// Example:
//  m := map[string]interface{}{
//  	"foo": map[string]interface{}{
//  		"bar": "baz",
//  	},
//  }
//
//  val, ok := Dig(m, "foo.bar")
//  // val => "baz", ok => true
//
func DigB(root interface{}, path string) (interface{}, bool) {
	if val, err := DigE(root, path); err == nil {
		return val, true
	}

	return nil, false
}

//
// DigE returns the value at a given path or nil as first result, and an error as second.
// Use the Is* functions to further check error type.
//
// Example:
//  m := map[string]interface{}{
//  	"foo": map[string]interface{}{
//  		"bar": "baz",
//  	},
//  }
//
//  val, err := DigE(m, "foo.baz")
//  // val => nil, IsNotFoundError(err) => true
//
func DigE(root interface{}, path string) (interface{}, error) {
	arr, err := Parse(path)
	if err != nil {
		return nil, err
	}

	return DigArrE(root, arr)
}

//
// DigArr does the same as Dig but instead of a string key, it expects an array of keys which are not further parsed.
//
func DigArr(root interface{}, arr []interface{}) interface{} {
	val, _ := DigArrE(root, arr)
	return val
}

//
// DigArrB does the same as DigB but instead of a string key, it expects an array of keys which are not further parsed.
//
// Example:
//  m := map[string]interface{}{
//  	"foo": map[string]interface{}{
//  		"bar": "baz",
//  	},
//  }
//
//  val, ok := DigArr(m, []interface{}{"foo", "bar"})
//  // val => "baz", ok => true
//
func DigArrB(root interface{}, arr []interface{}) (interface{}, bool) {
	if val, err := DigArrE(root, arr); err == nil {
		return val, true
	}

	return nil, false
}

//
// DigArrE does the same as DigArrB but returns an error as second result.
//
// Example:
//  m := map[string]interface{}{
//  	"foo": map[string]interface{}{
//  		"bar": "baz",
//  	},
//  }
//
//  val, ok := DigArrE(m, []interface{}{"foo", "baz"})
//  // val => nil, IsNotFoundError(err) => true
//
func DigArrE(root interface{}, arr []interface{}) (interface{}, error) {
	top := interface{}(root)
	curPath := ""

	for _, key := range arr {
		switch key.(type) {
		case string:
			k := key.(string)

			if lvl, ok := top.(map[string]interface{}); ok {
				if val, ok := lvl[k]; ok {
					top = val
				} else {
					return nil, NotFoundError(curPath)
				}
			} else if lvl, ok := top.(map[interface{}]interface{}); ok {
				if val, ok := lvl[k]; ok {
					top = val
				} else {
					return nil, NotFoundError(curPath)
				}
			} else {
				return nil, NotMapError(curPath)
			}

			curPath = fmt.Sprintf("%s.%s", curPath, k)
		case int:
			k := key.(int)

			if lvl, ok := top.([]interface{}); ok {
				if k < 0 {
					k = len(lvl) + k
				}

				if k < 0 || k >= len(lvl) {
					return nil, NotFoundError(curPath)
				}

				top = lvl[k]
			} else {
				return nil, NotArrayError(curPath)
			}

			curPath = fmt.Sprintf("%s.%d", curPath, k)
		}
	}

	return top, nil
}
