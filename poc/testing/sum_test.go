package testing

import (
	"testing"
)

func TestSumWithValidArgs(t *testing.T) {
	res := sumTwoDigits(10, 20)
	if res == 30 {
		t.Logf("%v + %v = %v", 10, 20, res)
	} else {
		t.Errorf("excepted output is 30 but we get %v ", res)
	}
}

func TestSumWithInvalidArgs(t *testing.T) {
	res := sumTwoDigits(20, 30)
	if res == 30 {
		t.Logf("%v + %v = %v", 10, 20, res)
	} else {
		t.Errorf("excepted output is 30 but we get %v ", res)
	}
}
