package bloom

import (
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hashHelpers"
)

type Bloom struct {
	Hashes  uint64
	Buckets uint64
	m       map[uint64]bool
}

func OptimalBuckets(n uint64, p float64) uint64 {
	return uint64(-(float64(n) * math.Log(p)) / (math.Log(2) * math.Log(2)))
}

// See https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func OptimalHashFunctions(p float64) uint64 {
	return uint64(-1.44 * math.Log2(p))
}

func New(n uint64, p float64) Bloom {
	buckets := OptimalBuckets(n, p)
	hashes := OptimalHashFunctions(p)

	return Bloom{
		Hashes:  hashes,
		Buckets: buckets,
		m:       make(map[uint64]bool, buckets),
	}
}

func (b Bloom) EstimateSize() float64 {
	x := 0.0
	for _, v := range b.m {
		if v {
			x++
		}
	}
	return -(float64(b.Buckets) / float64(b.Hashes)) *
		math.Log(1.0-(x/float64(b.Buckets)))
}

func (b Bloom) Add(data []byte) {
	hashes := hashHelpers.GetHashes(b.Hashes, data)
	for _, h := range hashes {
		b.m[h%b.Buckets] = true
	}
}

func (b Bloom) Contains(data []byte) bool {
	found := true
	hashes := hashHelpers.GetHashes(b.Hashes, data)
	for _, h := range hashes {
		if !b.m[h%b.Buckets] {
			found = false
		}
	}
	return found
}
