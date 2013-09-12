package quiz

import (
	"strings"
	"testing"
)

func newHarness() *testHelperHarness {
	return &testHelperHarness{false, ""}
}

type testHelperHarness struct {
	Failed  bool
	Message string
}

func (t *testHelperHarness) Fail() {
	t.Failed = true
}

func (t *testHelperHarness) FailNow() {
	t.Failed = true
}

func (t *testHelperHarness) Log(line string) {
	t.Message += line
}

func (t *testHelperHarness) Expect(target interface{}) *Expectation {
	return &Expectation{t: t, target: target}
}

func TestToEqual(t *testing.T) {
	h := newHarness()

	h.Expect(2).ToEqual(2)

	if h.Failed {
		t.Fail()
	}
}

func TestToEqualFail(t *testing.T) {
	h := newHarness()

	h.Expect(1).ToEqual(2)

	if !strings.Contains(h.Message, "Expected 2 to equal 1.") {
		t.Fail()
	}
}

func TestToBeFalse(t *testing.T) {
	h := newHarness()

	h.Expect(1 == 2).ToBeFalse()

	if h.Failed {
		t.Fail()
	}
}

func TestToBeFalseFail(t *testing.T) {
	h := newHarness()

	h.Expect(2 == 2).ToBeFalse()

	if !strings.Contains(h.Message, "Expected true to be false") {
		t.Fail()
	}
}

func TestToBeTrue(t *testing.T) {
	h := newHarness()

	h.Expect(1 == 1).ToBeTrue()

	if h.Failed {
		t.Fail()
	}
}

func TestToBeTrueFail(t *testing.T) {
	h := newHarness()

	h.Expect(2 == 3).ToBeTrue()

	if !strings.Contains(h.Message, "Expected false to be true") {
		t.Fail()
	}
}

func TestToBeGreaterThan(t *testing.T) {
	h := newHarness()

	h.Expect(1).ToBeGreaterThan(0)

	if h.Failed {
		t.Fail()
	}
}

func TestToBeGreaterThanFail(t *testing.T) {
	h := newHarness()

	h.Expect(0).ToBeGreaterThan(1)

	if !strings.Contains(h.Message, "Expected 0 to be greater than 1") {
		t.Fail()
	}
}

func TestToBeLessThan(t *testing.T) {
	h := newHarness()

	h.Expect(1).ToBeLessThan(2)

	if h.Failed {
		t.Fail()
	}
}

func TestToContain(t *testing.T) {
	h := newHarness()

	h.Expect("Hello world").ToContain("world")

	if h.Failed {
		t.Fail()
	}
}

func TestToContainFail(t *testing.T) {
	h := newHarness()

	h.Expect("Hello world").ToContain("Goodbye")

	if !strings.Contains(h.Message, "Expected Hello world to contain Goodbye") {
		t.Fail()
	}
}

func TestArrayToContain(t *testing.T) {
	h := newHarness()

	h.Expect([]string{"a"}).ToContain("a")

	if h.Failed {
		t.Fail()
	}
}
