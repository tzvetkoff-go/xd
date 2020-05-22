package xd_test

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
)

// Compare ...
func Compare(v1 interface{}, v2 interface{}) bool {
	return reflect.DeepEqual(v1, v2)
}

// Here ...
func Here() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}
