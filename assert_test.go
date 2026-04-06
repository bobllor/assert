package assert

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	failString = "test assertion failed, expected failure but got passing"
)

type MockT struct {
	Errors []string
}

// Contains checks if a substring is found in any of the strings
// inside MockT.Errors.
// If one is found, then return true, otherwise it will return false.
func (m *MockT) Contains(substring string) bool {
	for _, str := range m.Errors {
		if strings.Contains(str, substring) {
			return true
		}
	}

	return false
}

// IsEmpty returns a bool if MockT.Errors is empty.
// It will return true if it is empty, otherwise it will return false.
func (m *MockT) IsEmpty() bool {
	return len(m.Errors) == 0
}

func (m *MockT) Fatal(args ...any) {
	m.Errors = append(m.Errors, fmt.Sprint(args...))
}

func (m *MockT) Fatalf(format string, args ...any) {
	m.Errors = append(m.Errors, fmt.Sprintf(format, args...))
}

func TestAssertNil(t *testing.T) {
	_, err := os.ReadDir(t.TempDir())
	Nil(t, err)
	Nil(t, nil)
}

func TestAssertNilFail(t *testing.T) {
	mt := &MockT{}

	_, err := os.ReadDir("pathdoesnotexist/yes/no/maybeso")

	Nil(mt, err)
	Nil(mt, "string")

	if !mt.Contains(errNilFail) {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertNilSpecial(t *testing.T) {
	var pp *int
	var ch chan string
	var slice []string
	var mp map[string]string
	var intf any
	var fn func()

	Nil(t, pp)
	Nil(t, ch)
	Nil(t, slice)
	Nil(t, mp)
	Nil(t, intf)
	Nil(t, fn)
}

func TestAssertNilPtr(t *testing.T) {
	var ch *chan string
	var slice *[]string
	var mp *map[string]string
	var intf *any
	var fn *func()

	Nil(t, ch)
	Nil(t, slice)
	Nil(t, mp)
	Nil(t, intf)
	Nil(t, fn)
}

func TestAssertNilStruct(t *testing.T) {
	type Temp struct {
		Time *time.Time
	}

	v := Temp{}

	Nil(t, v.Time)
}

func TestAssertNilPtrFail(t *testing.T) {
	mt := &MockT{}
	i := 15
	date := time.Now().UTC()
	str := "string"

	Nil(mt, &i)
	Nil(mt, &date)
	Nil(mt, mt)
	Nil(mt, &str)

	if !mt.Contains(errNilFail) {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertNotNil(t *testing.T) {
	_, err := os.ReadDir("pathdoesnotexist/yes/no/maybeso")
	NotNil(t, err)

	var someVar int
	ch := make(chan string)

	NotNil(t, someVar)
	NotNil(t, ch)
	NotNil(t, "string")
	NotNil(t, 1+1)
}

func TestAssertNotNilPtr(t *testing.T) {
	type Temp struct {
		Time *time.Time
	}

	now := time.Now().UTC()
	v := Temp{
		Time: &now,
	}
	var temp Temp

	NotNil(t, v.Time)
	NotNil(t, temp)
}

func TestAssertTrue(t *testing.T) {
	base := "this is a test sentence"
	True(t, strings.Contains(base, "sentence"))
}

func TestAssertTrueFail(t *testing.T) {
	mt := &MockT{}

	True(mt, false)

	if !mt.Contains(errTrueFail) {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertFalse(t *testing.T) {
	base := "tester"
	False(t, strings.Contains(base, "easter"))
}

func TestAssertFalseFail(t *testing.T) {
	mt := &MockT{}

	False(mt, true)

	if !mt.Contains(errFalseFail) {
		t.Fatal(failString, mt.Errors)
	}
}
