package quiz

import (
	"strings"
	"testing"
)

func TestTo(t *testing.T) {
	e := NewExpectation(&spyTestHarness{}, "foo")
	e.To.Equal("foo")

	if e.To.failure == true {
		t.Log("Expected foo == foo")
		t.Fail()
	}

	e.To.Equal("bar")

	if e.To.failure == false {
		t.Log("Expected failure since foo != bar")
		t.Fail()
	}
}

func TestToNot(t *testing.T) {
	e := NewExpectation(&spyTestHarness{}, "foo")
	e.ToNot.Equal("bar")

	if e.ToNot.failure == true {
		t.Log("Expected foo != bar")
		t.Fail()
	}

	e.ToNot.Equal("foo")

	if e.ToNot.failure == false {
		t.Log("Expected failure since foo == foo")
		t.Fail()
	}
}

// Deprecated

func TestToEqual(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(2).ToEqual(2)

	if h.Failed {
		t.Fail()
	}
}

func TestToEqualFail(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(1).ToEqual(2)

	if !strings.Contains(h.Message, "Expected 1 to equal 2.") {
		t.Log("'"+h.Message+"' does not contain", "'Expect 1 to equal 2.'")
		t.Fail()
	}
}

func TestToBeFalse(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(1 == 2).ToBeFalse()

	if h.Failed {
		t.Fail()
	}
}

func TestToBeFalseFail(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(2 == 2).ToBeFalse()

	if !strings.Contains(h.Message, "Expected true to be false") {
		t.Fail()
	}
}

func TestToBeTrue(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(1 == 1).ToBeTrue()

	if h.Failed {
		t.Fail()
	}
}

func TestToBeTrueFail(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(2 == 3).ToBeTrue()

	if !strings.Contains(h.Message, "Expected false to be true") {
		t.Fail()
	}
}

func TestToBeGreaterThan(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(1).ToBeGreaterThan(0)

	if h.Failed {
		t.Fail()
	}
}

func TestToBeGreaterThanFail(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(0).ToBeGreaterThan(1)

	if !strings.Contains(h.Message, "Expected 0 to be greater than 1") {
		t.Fail()
	}
}

func TestToBeLessThan(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect(1).ToBeLessThan(2)

	if h.Failed {
		t.Fail()
	}
}

func TestToContain(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect("Hello world").ToContain("world")

	if h.Failed {
		t.Fail()
	}
}

func TestToContainFail(t *testing.T) {
	h := newSpyTestHarness()

	h.Expect("Hello world").ToContain("Goodbye")

	if !strings.Contains(h.Message, "Expected Hello world to contain Goodbye") {
		t.Fail()
	}
}
