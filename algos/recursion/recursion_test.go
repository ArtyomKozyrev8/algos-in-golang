package recursion

import (
	"testing"
)

func TestFibN(t *testing.T) {
	res := FibN(19)
	if res != 4181 {
		t.Errorf("%d != 4181", res)
	}

	res = FibN(12)
	if res != 144 {
		t.Errorf("%d != 144", res)
	}

	res = FibN(1)
	if res != 1 {
		t.Errorf("%d != 1", res)
	}

	res = FibN(2)
	if res != 1 {
		t.Errorf("%d != 1", res)
	}

	res = FibN(0)
	if res != 0 {
		t.Errorf("%d != 0", res)
	}

}
