package generator

import "testing"

func TestUnique(t *testing.T) {
	size := 1000000
	a := RandomStrings(size)
	b := RandomStrings(size)

	for i := 0; i < size; i++ {
		if a[i] == b[i] {
			t.Errorf("a[i] == b[i], %v == %v\n", a[i], b[i])
		}
	}

}
