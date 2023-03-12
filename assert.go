package pkgo

import (
	"fmt"
)

var _AssertionEnabled = true

func SetAssertion(enabled bool) {
	_AssertionEnabled = enabled
}

func AssertNil(v any, msg string) {
	if _AssertionEnabled && v != nil {
		panic(fmt.Sprintf("assertion failed[is nil]: %s", msg))
	}
}

func AssertNotNil(v any, msg string) {
	if _AssertionEnabled && v == nil {
		panic(fmt.Sprintf("assertion failed[not nil]: %s", msg))
	}
}

func AssertHasText(v string, msg string) {
	if _AssertionEnabled && v == "" {
		panic(fmt.Sprintf("assertion failed[has text]: %s", msg))
	}
}

func AssertEmpty(v string, msg string) {
	if _AssertionEnabled && v != "" {
		panic(fmt.Sprintf("assertion failed[is empty]: %s", msg))
	}
}

func AssertTrue(v bool, msg string) {
	if _AssertionEnabled && v != true {
		panic(fmt.Sprintf("assertion failed[is true]: %s", msg))
	}
}
