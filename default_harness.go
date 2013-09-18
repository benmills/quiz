package quiz

import (
	"fmt"
	"testing"
)

type defaultHarness struct {
	t *testing.T
}

func (harness defaultHarness) FailNow() {
	harness.t.FailNow()
}

func (harness defaultHarness) Log(line string) {
	fmt.Printf(line)
}

func (harness defaultHarness) Expect(target interface{}) *Expectation {
	return NewExpectation(harness, target)
}
