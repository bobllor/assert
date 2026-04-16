package assert

import (
	"os"
	"strings"
	"testing"
	"time"
)

const (
	failString = "test assertion failed, expected failure but got passing"
)

func TestAssertNil(t *testing.T) {
	_, err := os.ReadDir(t.TempDir())
	Nil(t, err)
	Nil(t, nil)
}

func TestAssertNilFail(t *testing.T) {
	mt := NewRecorder()

	_, err := os.ReadDir("pathdoesnotexist/yes/no/maybeso")

	Nil(mt, err)
	Nil(mt, "string")

	if !mt.Contains("is not nil") {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertNilAll(t *testing.T) {
	var pp *int
	var xd *string
	var lp *struct{}

	NilAll(t, pp, xd, lp)
}

func TestAssertNilAllFail(t *testing.T) {
	mt := NewRecorder()
	NilAll(mt, "some string value here")

	if !mt.Contains("is not nil") {
		t.Fatal(failString, mt.Errors)

	}
}

func TestAssertNilSpecialCases(t *testing.T) {
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
	mt := NewRecorder()
	i := 15
	date := time.Now().UTC()
	str := "string"

	Nil(mt, &i)
	Nil(mt, &date)
	Nil(mt, mt)
	Nil(mt, &str)

	if !mt.Contains("is not nil") {
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

func TestAssertNotNilAll(t *testing.T) {
	NotNilAll(t, "string", 123, 15.3, 'H', true, false)
}

func TestAssertNotNilAllPointers(t *testing.T) {
	pp := 3
	xd := "yes"
	st := struct{}{}

	NotNilAll(t, &pp, &xd, &st)
}

func TestAssertNotNilAllFail(t *testing.T) {
	mt := NewRecorder()
	NotNilAll(mt, "fdsa", 123, nil)

	if !mt.Contains("is nil") {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertTrue(t *testing.T) {
	base := "this is a test sentence"
	True(t, strings.Contains(base, "sentence"))
}

func TestAssertTrueAll(t *testing.T) {
	var z *int
	TrueAll(t, 10 == 100/10, true, strings.Contains("test sentence", "test"), z == nil)
}

func TestAssertTrueFail(t *testing.T) {
	mt := NewRecorder()

	conditions := []bool{
		false,
		1 != 2-1,
		strings.Contains("sentence", "z"),
		len([]byte{1, 2, 3}) == 2,
	}

	for _, cond := range conditions {
		True(mt, cond)
	}

	Equal(t, len(mt.Errors), len(conditions))
	if !mt.ContainsAllErrors("false condition") {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertFalse(t *testing.T) {
	base := "tester"
	False(t, strings.Contains(base, "easter"))
}

func TestAssertFalseFail(t *testing.T) {
	mt := NewRecorder()

	False(mt, true)

	if !mt.Contains("true condition") {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertContains(t *testing.T) {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	substr := "consectetur adipiscing"

	Contains(t, str, substr)
}

func TestAssertContainsFail(t *testing.T) {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	substr := "xxx"

	mt := NewRecorder()

	Contains(mt, str, substr)

	if !mt.Contains("not found in") {
		t.Fatal(failString, mt.Errors)
	}
}

func TestAssertContainsAny(t *testing.T) {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	substrings := []string{
		"fdsa",
		"h11",
		"432",
		"3333",
		"amet",
	}

	ContainsAny(t, str, substrings...)
}

func TestAssertNotContains(t *testing.T) {
	str := "Lorem ipsum"
	substr := "Curious"

	NotContains(t, str, substr)
}

func TestAssertNotContainsAny(t *testing.T) {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	substrings := []string{
		"fdsa",
		"h11",
		"432",
		"3333",
		"George",
	}

	NotContainsAny(t, str, substrings...)
}
