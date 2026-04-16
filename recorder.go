package assert

import (
	"fmt"
	"strings"
)

// Recorder is used for recording fatal output during an assertion
// failure.
type Recorder struct {
	Errors []string
}

// NewRecorder creates a new Recorder struct for use.
func NewRecorder() *Recorder {
	return &Recorder{}
}

// Contains checks if a substring is found in any of the strings
// inside t.Errors.
//
// If one is found, then return true, otherwise it will return false.
func (r *Recorder) Contains(substring string) bool {
	for _, str := range r.Errors {
		if strings.Contains(str, substring) {
			return true
		}
	}

	return false
}

// ContainsAllErrors checks if the substring can be found in all strings
// of t.Errors.
//
// It will return true if all values of t.Errors contains the substring,
// otherwise it will return false.
func (r *Recorder) ContainsAllErrors(substring string) bool {
	for _, str := range r.Errors {
		if !strings.Contains(str, substring) {
			return false
		}
	}

	return true
}

// IsEmpty returns a true if t.Errors is empty.
func (r *Recorder) IsEmpty() bool {
	return len(r.Errors) == 0
}

// Fatal appends the args into t.Errors.
func (r *Recorder) Fatal(args ...any) {
	r.Errors = append(r.Errors, fmt.Sprint(args...))
}

// Fatalf appends the formatted args into t.Errors.
func (r *Recorder) Fatalf(format string, args ...any) {
	r.Errors = append(r.Errors, fmt.Sprintf(format, args...))
}
