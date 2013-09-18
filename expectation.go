package quiz

type Expectation struct {
	t      TestHarness
	target interface{}
	To     *Matcher
	ToNot  *Matcher
}

func NewExpectation(t TestHarness, target interface{}) *Expectation {
	return &Expectation{
		t:      t,
		target: target,
		To:     PositiveMatcher(t, target),
		ToNot:  NegativeMatcher(t, target),
	}
}

// Deprecated

func (expect *Expectation) ToEqual(value interface{}) {
	expect.To.Equal(value)
}

func (expect *Expectation) ToBeTrue() {
	expect.To.BeTrue()
}

func (expect *Expectation) ToBeFalse() {
	expect.To.BeFalse()
}

func (expect *Expectation) ToBeGreaterThan(value int) {
	expect.To.BeGreaterThan(value)
}

func (expect *Expectation) ToBeLessThan(value int) {
	expect.To.BeLessThan(value)
}

func (expect *Expectation) ToContain(value interface{}) {
	expect.To.Contain(value)
}
