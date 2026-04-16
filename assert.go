package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Tester interface {
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

// Equal asserts if two values are equal to each other.
func Equal(t Tester, v any, compare any) {
	equal := reflect.DeepEqual(v, compare)
	callerInfo := getCallerInfo()

	if !equal {
		t.Fatalf("%s: values are not equal (expected %v, got %v)", callerInfo, compare, v)
	}
}

// NotEqual asserts if two values are not equal to each other.
func NotEqual(t Tester, v any, compare any) {
	equal := reflect.DeepEqual(v, compare)
	callerInfo := getCallerInfo()

	if equal {
		t.Fatalf("%s: values are equal (%v == %v)", callerInfo, v, compare)
	}
}

// Nil asserts if the value is nil.
func Nil(t Tester, v any) {
	callerInfo := getCallerInfo()

	if !checkNil(v) {
		t.Fatalf("%s: value is not nil (%v)", callerInfo, v)
	}
}

// NilAll asserts if all given values are nil.
func NilAll(t Tester, vs ...any) {
	callerInfo := getCallerInfo()

	for _, v := range vs {
		if !checkNil(v) {
			t.Fatalf("%s: value is not nil (%v in %v)", callerInfo, v, vs)
		}
	}
}

// NotNil asserts if the value is not nil.
func NotNil(t Tester, v any) {
	callerInfo := getCallerInfo()

	if checkNil(v) {
		t.Fatalf("%s: value is nil", callerInfo, v)
	}
}

// NotNilAll asserts if all the given values are not nil.
func NotNilAll(t Tester, vs ...any) {
	callerInfo := getCallerInfo()

	for _, v := range vs {
		if checkNil(v) {
			t.Fatalf("%s: value is nil (%v in %v)", callerInfo, v, vs)
		}
	}
}

// True asserts if the condition is true.
func True(t Tester, cond bool) {
	callerInfo := getCallerInfo()

	if !cond {
		t.Fatalf("%s: got false condition", callerInfo)
	}
}

// TrueAll asserts if all conditions are true.
func TrueAll(t Tester, conds ...bool) {
	callerInfo := getCallerInfo()

	for _, cond := range conds {
		if !cond {
			t.Fatalf("%s: got false condition (%v)", callerInfo, conds)
		}
	}
}

// False asserts if the condition is false.
func False(t Tester, cond bool) {
	callerInfo := getCallerInfo()

	if cond {
		t.Fatalf("%s: got true condition", callerInfo)
	}
}

// FalseAll asserts if all conditions are false.
func FalseAll(t Tester, conds ...bool) {
	callerInfo := getCallerInfo()

	for _, cond := range conds {
		if cond {
			t.Fatalf("%s: got true condition (%v)", callerInfo, conds)
		}
	}
}

// Contains asserts if a string contains a substring.
func Contains(t Tester, s string, substr string) {
	callerInfo := getCallerInfo()

	if !strings.Contains(s, substr) {
		t.Fatalf("%s: substring '%s' not found in '%s'", callerInfo, substr, s)
	}
}

// NotContains asserts if a string does not contain a substring.
func NotContains(t Tester, s string, substr string) {
	callerInfo := getCallerInfo()

	if strings.Contains(s, substr) {
		t.Fatalf("%s: substring '%s' found in '%s'", callerInfo, substr, s)
	}
}

// ContainsAny asserts if any substring in substrings is found in the string.
func ContainsAny(t Tester, s string, substrings ...string) {
	callerInfo := getCallerInfo()

	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return
		}
	}

	t.Fatalf("%s: no substrings found in '%s'", callerInfo, s)
}

// NotContainsAny asserts if any substring in substrings is not found in the string.
func NotContainsAny(t Tester, s string, substrings ...string) {
	callerInfo := getCallerInfo()

	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			t.Fatalf("%s: substring '%s' found in '%s'", callerInfo, substr, s)
		}
	}

}

// checkNil checks if the value is nil.
// It will return true if it is nil, otherwise return false.
func checkNil(val any) bool {
	v := reflect.ValueOf(val)

	// val is a literal nil (not interface)
	if v.Kind() == reflect.Invalid {
		if val == nil {
			return true
		}
	} else {
		switch v.Type().Kind() {
		case reflect.Ptr:
			elem := v.Elem()

			if elem.Kind() == reflect.Invalid {
				return true
			}
		// pointer not included due to potential struct
		// pointers
		case reflect.Chan, reflect.Slice,
			reflect.Map, reflect.Interface, reflect.Func:
			if v.IsNil() {
				return true
			}
		default:
			if val == nil {
				return true
			}
		}
	}

	return false
}

// getCallerInfo returns the string containing the caller metadata in the format:
// "<test file>:<line> (<function name> assertion failed)"
// The metadata is to the main parent caller outside of assert.
//
// If the attempt in getting the caller information fails, then it will return
// the string "caller retrieval could not be completed".
func getCallerInfo() (callerStr string) {
	// parent -> assert -> getCallerInfo (2 levels)
	skipLevel := 2
	pc, file, n, ok := runtime.Caller(skipLevel)

	if !ok {
		callerStr = "caller retrieval could not be completed"
	} else {
		parentCaller := filepath.Base(runtime.FuncForPC(pc).Name())
		fileName := filepath.Base(file)

		callerStr = fmt.Sprintf("%s:%d (%s assertion failed)", fileName, n, parentCaller)
	}

	return callerStr
}
