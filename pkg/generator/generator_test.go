package generator

import "testing"

func TestUnique(t *testing.T) {
	a := RandomStrings(10)
	b := RandomStrings(10)

	for i := 0; i < 10; i++ {
		if a[i] == b[i] {
			t.Errorf("a[i] == b[i], %v == %v\n", a[i], b[i])
		}
	}

}
