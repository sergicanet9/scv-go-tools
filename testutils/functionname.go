package testutils

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// FunctionName returns the name of the given function as string
func FunctionName(t *testing.T, function interface{}) string {
	t.Helper()

	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}
