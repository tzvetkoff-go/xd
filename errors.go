package xd

import (
	"fmt"
	"strings"
)

// Error codes.
const (
	ErrorCodeParse    = iota // Path cannot be parsed
	ErrorCodeNotFound        // Element at current path does not exist
	ErrorCodeNotMap          // Element at current path is not a `map[string]interface{}`
	ErrorCodeNotArray        // Element at current path is not a `[]interface{}`
)

// Error represents an error returned by the functions in this package.
type Error struct {
	ErrorCode int
	Message   string
	Path      string
}

// Error implements the builtin error interface.
func (e *Error) Error() string {
	return e.Message
}

// ParseError returns an Error with ErrorCode set to ErrorCodeParse.
func ParseError(message string) error {
	return &Error{
		ErrorCode: ErrorCodeParse,
		Message:   message,
	}
}

// IsParseError tests if the error has ErrorCode set to ErrorCodeParse.
func IsParseError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.ErrorCode == ErrorCodeParse
	}

	return false
}

// NotFoundError returns an Error with ErrorCode set to ErrorCodeNotFound.
func NotFoundError(path string) error {
	if strings.HasPrefix(path, ".") {
		path = path[1:]
	}

	return &Error{
		ErrorCode: ErrorCodeNotFound,
		Message:   fmt.Sprintf("element at path `%s` not found", path),
		Path:      path,
	}
}

// IsNotFoundError tests if the error has ErrorCode set to ErrorCodeNotFound.
func IsNotFoundError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.ErrorCode == ErrorCodeNotFound
	}

	return false
}

// NotMapError returns an Error with ErrorCode set to ErrorCodeNotMap.
func NotMapError(path string) error {
	if strings.HasPrefix(path, ".") {
		path = path[1:]
	}

	return &Error{
		ErrorCode: ErrorCodeNotMap,
		Message:   fmt.Sprintf("element at path `%s` not of type map[string]interface{}", path),
		Path:      path,
	}
}

// IsNotMapError tests if the error has ErrorCode set to ErrorCodeNotMap.
func IsNotMapError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.ErrorCode == ErrorCodeNotMap
	}

	return false
}

// NotArrayError returns an Error with ErrorCode set to ErrorCodeNotArray.
func NotArrayError(path string) error {
	if strings.HasPrefix(path, ".") {
		path = path[1:]
	}

	return &Error{
		ErrorCode: ErrorCodeNotArray,
		Message:   fmt.Sprintf("element at path `%s` not of type []interface{}", path),
		Path:      path,
	}
}

// IsNotArrayError tests if the error has ErrorCode set to ErrorCodeNotArray.
func IsNotArrayError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.ErrorCode == ErrorCodeNotArray
	}

	return false
}
