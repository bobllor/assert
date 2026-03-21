package assert

import (
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

// NotNil asserts if the value is not nil.
func NotNil(t *testing.T, val any) {
	if val == nil {
		t.Fatalf("%s: val is nil", assertionErrStr)
	}
}

// Nil asserts if the value is nil.
func Nil(t *testing.T, val any) {
	if val != nil {
		t.Fatalf("%s: val is not nil (%v)", assertionErrStr, val)
	}
}
