package quiz

import (
	"testing"
)

func TestQuiz(t *testing.T) {
	test := Test(t)

	test.Expect("foo").ToEqual("foo")
	test.Expect([]string{"bar", "baz"}).To.Contain("baz")
	test.Expect(true).ToNot.BeFalse()
}
