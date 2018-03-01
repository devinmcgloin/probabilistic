package hh

import (
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestMultiUnique(t *testing.T) {
	values := []string{"", "t", "s", "c", "sdsdjksd", "sdsdS", "sdsfsdfs", "hi2ounx"}
	values = append(values, generator.RandomStrings(6000)...)

	unique := map[uint64]int{}
	for _, v := range values {
		hashes := GetHashes(55, []byte(v))
		for _, hash := range hashes {
			if unique[hash] == 2 {
				t.Errorf("On input %v hash %v occurs with 55 hash functions\n", v, hash)
			} else {
				v, ok := unique[hash]
				if ok {
					unique[hash] = v + 1
				} else {
					unique[hash] = 1
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
			h := FNVBias([]byte(v), i)
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
