package hashHelpers

import "testing"

func TestUnique(t *testing.T) {
	values := []string{"t", "s", "c", "sdsdjksd", "sdsdS", "sdsfsdfs", "hi2ounx"}

	for _, v := range values {
		for i := uint64(2); i < 12; i++ {
			unique := map[uint64]bool{}
			hashes := GetHashes(i, []byte(v))
			for _, hash := range hashes {
				if unique[hash] {
					t.Errorf("On input %v hash %v occurs with %d hash functions\n", v, hashes, i)
				} else {
					unique[hash] = true
				}
			}
		}
	}

}
