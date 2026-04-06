package assert

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestAssertNil(t *testing.T) {
	_, err := os.ReadDir(t.TempDir())
	Nil(t, err)
}

func TestAssertNilPtr(t *testing.T) {
	type Temp struct {
		Time *time.Time
	}

	v := Temp{}

	var pp *int

	Nil(t, v.Time)
	Nil(t, pp)
}

func TestAssertNotNil(t *testing.T) {
	_, err := os.ReadDir("pathdoesnotexist/yes/no/maybeso")
	NotNil(t, err)

	var someVar int
	ch := make(chan string)

	NotNil(t, someVar)
	NotNil(t, ch)
}

func TestAssertNotNilPtr(t *testing.T) {
	type Temp struct {
		Time *time.Time
	}

	now := time.Now().UTC()
	v := Temp{
		Time: &now,
	}

	NotNil(t, v.Time)
}

func TestAssertTrue(t *testing.T) {
	base := "this is a test sentence"
	True(t, strings.Contains(base, "sentence"))
}

func TestAssertFalse(t *testing.T) {
	base := "tester"
	False(t, strings.Contains(base, "easter"))
}
