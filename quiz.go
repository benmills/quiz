package quiz

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func Test(t *testing.T) *defaultHarness {
	return &defaultHarness{t}
}

func NewExpectation(t TestHarness, target interface{}) *Expectation {
	return &Expectation{t: t, target: target}
}

type TestHarness interface {
	Fail()
	Log(string)
	Expect(interface{}) *Expectation
}

type defaultHarness struct {
	t *testing.T
}

func (harness defaultHarness) Fail() {
	harness.t.Fail()
}

func (harness defaultHarness) Log(line string) {
	fmt.Printf(line)
}

func (harness defaultHarness) Expect(target interface{}) *Expectation {
	return &Expectation{t: harness, target: target}
}

type Expectation struct {
	t      TestHarness
	target interface{}
}

type assertion struct {
	failure        bool
	failureMessage string
	messageParts   []interface{}
	expect         *Expectation
}

func (a assertion) eval(expect *Expectation) {
	if a.failure {
		_, file, line, _ := runtime.Caller(2)
		expect.t.Fail()
		expect.t.Log(fmt.Sprintf(a.failureMessage+"\n  %s:%d\n", append(a.messageParts, file, line)...))
	}
}

func (expect *Expectation) ToEqual(value interface{}) {
	assertion{
		failure:        expect.target != value,
		failureMessage: "Expected %v to equal %v.",
		messageParts:   []interface{}{value, expect.target},
	}.eval(expect)
}

func (expect *Expectation) ToBeTrue() {
	assertion{
		failure:        expect.target != true,
		failureMessage: "Expected %v to be true.",
		messageParts:   []interface{}{expect.target},
	}.eval(expect)
}

func (expect *Expectation) ToBeFalse() {
	assertion{
		failure:        expect.target != false,
		failureMessage: "Expected %v to be false.",
		messageParts:   []interface{}{expect.target},
	}.eval(expect)
}

func (expect *Expectation) ToBeLessThan(value int) {
	intTarget := expect.target.(int)
	assertion{
		failure:        intTarget > value,
		failureMessage: "Expected %v to be less than %v.",
		messageParts:   []interface{}{expect.target, value},
	}.eval(expect)
}

func (expect *Expectation) ToBeGreaterThan(value int) {
	intTarget := expect.target.(int)
	assertion{
		failure:        intTarget < value,
		failureMessage: "Expected %v to be greater than %v.",
		messageParts:   []interface{}{expect.target, value},
	}.eval(expect)
}

func (expect *Expectation) ToContain(value string) {
	stringTarget := expect.target.(string)
	assertion{
		failure:        !strings.Contains(stringTarget, value),
		failureMessage: "Expected %s to contain %s.",
		messageParts:   []interface{}{expect.target, value},
	}.eval(expect)
}
