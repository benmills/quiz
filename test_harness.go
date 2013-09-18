package quiz

type TestHarness interface {
	FailNow()
	Log(string)
	Expect(interface{}) *Expectation
}
