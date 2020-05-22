package xd

import (
	"strconv"
)

//
// Parse parses a string path and returns an array of keys.
// Keys can be either strings or integers.
// The format follows the common struct/array notation, for example:
//
//   Parse("foo.bar[0]")
//   // []interface{}{"foo", "bar", 0}
//
// The character "\" is used as an escape character in case a key contains "." or "[".
//
//   Parse("foo\\.bar")
//   // []interface{}{"foo.bar"}
//
func Parse(path string) ([]interface{}, error) {
	result := []interface{}{}

	keyString := ""

	prev := '\u0000'
	for _, char := range []rune(path) {
		switch char {
		case '\\':
			if prev == '\\' {
				keyString += "\\"
				prev = '🍺'
				continue
			}

			goto next
		case '.':
			if prev == '\\' {
				keyString += "."
				goto next
			}
			if prev == ']' || prev == '\u0000' {
				goto next
			}

			result = append(result, keyString)
			keyString = ""
		case '[':
			if prev == '\\' {
				keyString += "["
				goto next
			}
			if prev == ']' || prev == '\u0000' {
				goto next
			}

			result = append(result, keyString)
			keyString = ""
		case ']':
			i64, err := strconv.ParseInt(keyString, 0, 64)
			if err != nil {
				return nil, ParseError(err.Error())
			}

			result = append(result, int(i64))
			keyString = ""
		default:
			keyString += string(char)
		}

	next:
		prev = char
	}

	if prev == '\\' {
		return nil, ParseError("path ends with an escape character")
	}

	if keyString != "" || prev == '.' {
		result = append(result, keyString)
	}

	return result, nil
}
