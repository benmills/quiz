package quiz

import (
	"fmt"
	"strings"
)

type MatcherKind string

const (
	positive MatcherKind = "positive"
	negative MatcherKind = "negative"
)

type Matcher struct {
	kind    MatcherKind
	target  interface{}
	failure bool
	message string
	harness TestHarness
}

func PositiveMatcher(harness TestHarness, target interface{}) *Matcher {
	return &Matcher{kind: positive, target: target, harness: harness}
}

func NegativeMatcher(harness TestHarness, target interface{}) *Matcher {
	return &Matcher{kind: negative, target: target, harness: harness}
}

func (matcher *Matcher) Equal(value interface{}) *Matcher {
	matcher.eval(matcher.target != value, "Expected %v to equal %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeTrue() *Matcher {
	matcher.eval(matcher.target != true, "Expected %v to be true.", matcher.target)
	return matcher
}

func (matcher *Matcher) BeFalse() *Matcher {
	matcher.eval(matcher.target != false, "Expected %v to be false.", matcher.target)
	return matcher
}

func (matcher *Matcher) BeGreaterThan(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() < value, "Expected %v to be greater than %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) BeLessThan(value int) *Matcher {
	matcher.eval(matcher.targetAsInt() > value, "Expected %v to be less than %v.", matcher.target, value)
	return matcher
}

func (matcher *Matcher) Contain(value interface{}) *Matcher {
	matcher.eval(
		!strings.Contains(matcher.targetAsString(), toString(value)),
		"Expected %v to contain %v.", matcher.target, value,
	)

	return matcher
}

func (matcher *Matcher) targetAsInt() int {
	return matcher.target.(int)
}

func (matcher *Matcher) targetAsString() string {
	return toString(matcher.target)
}

func toString(value interface{}) string {
	return fmt.Sprint(value)
}

func (matcher *Matcher) eval(failure bool, message string, parts ...interface{}) {
	if matcher.kind == positive {
		matcher.failure = failure
	} else if matcher.kind == negative {
		matcher.failure = !failure
	}

	if matcher.failure {
		matcher.message = fmt.Sprintf(message, parts...)
		matcher.harness.Log(matcher.message)
		matcher.harness.FailNow()
	}
}
