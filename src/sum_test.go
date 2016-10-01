package webapp

import "testing"

type sumTest struct {
	in1, in2, out int
}

var sumTests = []sumTest{
	sumTest{1, 1, 2},
	sumTest{2, 2, 4},
	sumTest{100, 200, 300},
}

func TestSum(t *testing.T) {
	for _, dt := range sumTests {
		v := Sum(dt.in1, dt.in2)
		if v != dt.out {
			t.Errorf("Sum(%d, %d) = %d, want %d.", dt.in1, dt.in2, v, dt.out)
		}
	}
}
