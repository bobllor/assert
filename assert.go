package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
)

const (
	errTrueFail     = "value is false"
	errFalseFail    = "value is true"
	errEqualFail    = "does not equal"
	errNotEqualFail = "equals"
	errNilFail      = "value is not nil"
	errNotNilFail   = "value is nil"
)

type Tester interface {
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

// Equal asserts if two values are equal to each other.
func Equal(t Tester, v1 any, v2 any) {
	equal := reflect.DeepEqual(v1, v2)
	callerInfo := getCallerInfo()

	if !equal {
		t.Fatalf("%s: %v %s %v", callerInfo, v1, errEqualFail, v2)
	}
}

// NotEqual asserts if two values are not equal to each other.
func NotEqual(t Tester, v1 any, v2 any) {
	equal := reflect.DeepEqual(v1, v2)
	callerInfo := getCallerInfo()

	if equal {
		t.Fatalf("%s: %v %s %v", callerInfo, v1, errNotEqualFail, v2)
	}
}

// Nil asserts if the value is nil.
func Nil(t Tester, v any) {
	callerInfo := getCallerInfo()

	if !checkNil(v) {
		t.Fatalf("%s: %s (%v)", callerInfo, errNilFail, v)
	}
}

// NilAll asserts if all given values are nil.
func NilAll(t Tester, vs ...any) {
	callerInfo := getCallerInfo()

	for _, v := range vs {
		if !checkNil(v) {
			t.Fatalf("%s: %s (%v)", callerInfo, errNilFail, v)
		}
	}
}

// NotNil asserts if the value is not nil.
func NotNil(t Tester, v any) {
	callerInfo := getCallerInfo()

	if checkNil(v) {
		t.Fatalf("%s: %s (%v)", callerInfo, errNotNilFail, v)
	}
}

// NotNilAll asserts if all the given values are not nil.
func NotNilAll(t Tester, vs ...any) {
	callerInfo := getCallerInfo()

	for _, v := range vs {
		if checkNil(v) {
			t.Fatalf("%s: %s (%v)", callerInfo, errNotNilFail, v)
		}
	}
}

// True asserts if the condition is true.
func True(t Tester, cond bool) {
	callerInfo := getCallerInfo()

	if !cond {
		t.Fatalf("%s: %s", callerInfo, errTrueFail)
	}
}

// TrueAll asserts if all conditions are true.
func TrueAll(t Tester, conds ...bool) {
	callerInfo := getCallerInfo()

	for _, cond := range conds {
		if !cond {
			t.Fatalf("%s: %s", callerInfo, errTrueFail)
		}
	}
}

// False asserts if the condition is false.
func False(t Tester, cond bool) {
	callerInfo := getCallerInfo()

	if cond {
		t.Fatalf("%s: %s", callerInfo, errFalseFail)
	}
}

// FalseAll asserts if all conditions are false.
func FalseAll(t Tester, conds ...bool) {
	callerInfo := getCallerInfo()

	for _, cond := range conds {
		if cond {
			t.Fatalf("%s: %s", callerInfo, errFalseFail)
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
// "<file>:<line> (<function name> failed)"
// The metadata is to the main parent caller outside of assert.
//
// If the attempt in getting the caller information fails, then it will return
// the string "(caller info failed)".
func getCallerInfo() (callerStr string) {
	// parent -> assert -> getCallerInfo (2 levels)
	skipLevel := 2
	pc, file, n, ok := runtime.Caller(skipLevel)

	if !ok {
		callerStr = "(caller retrieval unexpected error)"
	} else {
		parentCaller := filepath.Base(runtime.FuncForPC(pc).Name())
		fileName := filepath.Base(file)

		callerStr = fmt.Sprintf("%s:%d (%s failed)", fileName, n, parentCaller)
	}

	return callerStr
}
