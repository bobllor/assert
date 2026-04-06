package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
)

const (
	errTrueFail     = "val is false"
	errFalseFail    = "val is true"
	errEqualFail    = "does not equal"
	errNotEqualFail = "equals"
	errNilFail      = "val is not nil"
	errNotNilFail   = "val is nil"
)

type Tester interface {
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

// Equal asserts if two values are equal to each other.
func Equal(t Tester, valOne any, valTwo any) {
	equal := reflect.DeepEqual(valOne, valTwo)
	callerInfo := getCallerInfo()

	if !equal {
		t.Fatalf("%s: %v %s %v", callerInfo, errEqualFail, valOne, valTwo)
	}
}

// NotEqual asserts if two values are not equal to each other.
func NotEqual(t Tester, valOne any, valTwo any) {
	equal := reflect.DeepEqual(valOne, valTwo)
	callerInfo := getCallerInfo()

	if equal {
		t.Fatalf("%s: %v %s %v", callerInfo, errNotEqualFail, valOne, valTwo)
	}
}

// Nil asserts if the value is nil.
func Nil(t Tester, val any) {
	callerInfo := getCallerInfo()

	if !checkNil(val) {
		t.Fatalf("%s: %s (%v)", callerInfo, errNilFail, val)
	}
}

// NotNil asserts if the value is not nil.
func NotNil(t Tester, val any) {
	callerInfo := getCallerInfo()

	if checkNil(val) {
		t.Fatalf("%s: %s (%v)", callerInfo, errNotNilFail, val)
	}
}

// True asserts if the value is true.
func True(t Tester, val bool) {
	callerInfo := getCallerInfo()

	if !val {
		t.Fatalf("%s: %s", callerInfo, errTrueFail)
	}
}

// False asserts if the value is false.
func False(t Tester, val bool) {
	callerInfo := getCallerInfo()

	if val {
		t.Fatalf("%s: %s", callerInfo, errFalseFail)
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
