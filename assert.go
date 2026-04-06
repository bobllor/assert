package assert

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	// assertionErrStr is a literal string "assertion error" used for formatting errors.
	assertionErrStr = "assertion failed"
)

// Equal asserts if two values are equal to each other.
func Equal(t *testing.T, valOne any, valTwo any) {
	equal := reflect.DeepEqual(valOne, valTwo)

	if !equal {
		t.Fatalf("%s: %v does not equal %v", assertionErrStr, valOne, valTwo)
	}
}

// NotEqual asserts if two values are not equal to each other.
func NotEqual(t *testing.T, valOne any, valTwo any) {
	equal := reflect.DeepEqual(valOne, valTwo)

	if equal {
		t.Fatalf("%s: %v equals %v", assertionErrStr, valOne, valTwo)
	}
}

// Nil asserts if the value is nil.
func Nil(t *testing.T, val any) {
	err := fmt.Errorf("%s: val is not nil (%v)", assertionErrStr, val)

	if !checkNil(val) {
		t.Fatal(err)
	}
}

// NotNil asserts if the value is not nil.
func NotNil(t *testing.T, val any) {
	err := fmt.Errorf("%s: val is nil (%v)", assertionErrStr, val)

	if checkNil(val) {
		t.Fatal(err)
	}
}

// True asserts if the value is true.
func True(t *testing.T, val bool) {
	if !val {
		t.Fatalf("%s: val is false", assertionErrStr)
	}
}

// False asserts if the value is false.
func False(t *testing.T, val bool) {
	if val {
		t.Fatalf("%s: val is true", assertionErrStr)
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
			if v.Elem().Kind() == reflect.Invalid {
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
