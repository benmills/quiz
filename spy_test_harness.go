package quiz

func newSpyTestHarness() *spyTestHarness {
	return &spyTestHarness{false, ""}
}

type spyTestHarness struct {
	Failed  bool
	Message string
}

func (t *spyTestHarness) Fail() {
	t.Failed = true
}

func (t *spyTestHarness) FailNow() {
	t.Failed = true
}

func (t *spyTestHarness) Log(line string) {
	t.Message += line
}

func (t *spyTestHarness) Expect(target interface{}) *Expectation {
	return NewExpectation(t, target)
}
