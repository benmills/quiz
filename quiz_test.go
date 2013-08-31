package quiz

import "testing"

func TestSimple(t *testing.T) {
	test := Test(t)

	test.Expect(1 == 2).ToBeFalse()
}

func TestSimpleTwo(t *testing.T) {
	test := Test(t)

	test.Expect(1).ToEqual(1)
	test.Expect(1 == 2).ToBeFalse()
	test.Expect(1 == 1).ToBeTrue()
	test.Expect(1).ToBeGreaterThan(0)
	test.Expect(1).ToBeLessThan(2)
	test.Expect("hello world").ToContain("world")
}
