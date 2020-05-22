//
// Package xd imlements functions for easy map[string]interface{} traversal.
//
// Example:
//  m := map[string]interface{}{
//  	"arr": []interface{}{
//  		"object",
//  		map[string]interface{}{
//  			"foo.bar": "baz",
//  		},
//  	},
//  }
//
//  val, ok := Dig(m, "arr[0].foo\\.bar")
//  // val => nil, ok => false
//
//  val, err := DigE(m, "arr[1].foo\\.bar")
//  // val => "baz", err => nil
//
package xd
