package quiz

import (
	"path/filepath"
	"runtime"
	"testing"
)

func testTrue(t *testing.T, ex bool) {
	if !ex {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("Failed at %s:%v\n", filepath.Base(file), line)
		t.Fail()
	}
}

func expectPositive(target interface{}) *Matcher {
	return PositiveMatcher(newSpyTestHarness(), target)
}

func expectNegative(target interface{}) *Matcher {
	return NegativeMatcher(newSpyTestHarness(), target)
}

func TestEqualSucceeds(t *testing.T) {
	testTrue(t, expectPositive("foo").Equal("foo").failure == false)
	testTrue(t, expectNegative("foo").Equal("bar").failure == false)
}

func TestEqualFails(t *testing.T) {
	failingMatch := expectPositive("foo").Equal("bar")
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected foo to equal bar.")
}

func TestBeTrueSucceeds(t *testing.T) {
	testTrue(t, expectPositive(true).BeTrue().failure == false)
	testTrue(t, expectNegative(false).BeTrue().failure == false)
}

func TestBeTrueFails(t *testing.T) {
	failingMatch := expectPositive(false).BeTrue()
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected false to be true.")
}

func TestBeFalseSucceeds(t *testing.T) {
	testTrue(t, expectPositive(false).BeFalse().failure == false)
	testTrue(t, expectNegative(false).BeFalse().failure == true)
}

func TestBeFalseFails(t *testing.T) {
	failingMatch := expectPositive(true).BeFalse()
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected true to be false.")
}

func TestBeGreaterThanSucceeds(t *testing.T) {
	testTrue(t, expectPositive(1).BeGreaterThan(0).failure == false)
	testTrue(t, expectNegative(1).BeGreaterThan(0).failure == true)
}

func TestBeGreaterThanFails(t *testing.T) {
	failingMatch := expectPositive(0).BeGreaterThan(1)
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected 0 to be greater than 1.")
}

func TestBeLessThanSucceeds(t *testing.T) {
	testTrue(t, expectPositive(0).BeLessThan(1).failure == false)
	testTrue(t, expectNegative(0).BeLessThan(1).failure == true)
}

func TestBeLessThanFails(t *testing.T) {
	failingMatch := expectPositive(1).BeLessThan(0)
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected 1 to be less than 0.")
}

func TestContainSucceeds(t *testing.T) {
	testTrue(t, expectPositive("a").Contain("a").failure == false)
	testTrue(t, expectNegative("a").Contain("a").failure == true)
}

func TestContainFails(t *testing.T) {
	failingMatch := expectPositive("a").Contain("b")
	testTrue(t, failingMatch.failure == true)
	testTrue(t, failingMatch.message == "Expected a to contain b.")
}

func TestContainWorksWithSlices(t *testing.T) {
	testTrue(t, expectPositive([]string{"a"}).Contain("a").failure == false)
}
