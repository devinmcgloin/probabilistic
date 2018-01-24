package hashHelpers

import (
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestMultiUnique(t *testing.T) {
	values := []string{"", "t", "s", "c", "sdsdjksd", "sdsdS", "sdsfsdfs", "hi2ounx"}
	values = append(values, generator.RandomStrings(1500)...)

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

func TestFNVUnique(t *testing.T) {
	values := []string{"", "t", "s", "c", "sdsdjksd", "sdsdS", "sdsfsdfs", "hi2ounx"}
	values = append(values, generator.RandomStrings(15000)...)

	for _, v := range values {

		unique := map[uint64]bool{}
		for i := uint64(0); i < 100; i++ {
			h := fnvBias([]byte(v), i)
			if unique[h] {
				t.Errorf("On input %v hash %v occurs with %d hash functions\n", v, h, i)
			} else {
				unique[h] = true
			}

		}
	}
}

var hash []uint64

func BenchmarkHash100(b *testing.B) {
	BenchHash(100, b)
}

func BenchmarkHash500(b *testing.B) {
	BenchHash(500, b)
}
func BenchmarkHash1500(b *testing.B) {
	BenchHash(1500, b)
}
func BenchmarkHash5000(b *testing.B) {
	BenchHash(5000, b)
}
func BenchHash(size int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		values := generator.RandomStrings(size)
		var h []uint64
		for _, v := range values {
			h = GetHashes(11, []byte(v))
		}
		hash = h
	}
}
