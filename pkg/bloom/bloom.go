package bloom

import (
	"errors"
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hh"
)

// Bloom represents a Bloom filter, including the number of hash functions used and size of backing
type Bloom struct {
	Hashes  uint64
	Buckets uint64
	m       []bool
}

// OptimalBuckets calculates the optimal number of buckets given the number of elements expected and desired threshold for false positives.
func OptimalBuckets(n uint64, p float64) uint64 {
	return uint64(-(float64(n) * math.Log(p)) / (math.Log(2) * math.Log(2)))
}

// OptimalHashFunctions determines the optimal number of hash functions to use for a desired
// false positive rate.
// See https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func OptimalHashFunctions(p float64) uint64 {
	return uint64(-1.44 * math.Log2(p))
}

// New constructs a new Bloom filter
func New(n uint64, p float64) Bloom {
	buckets := OptimalBuckets(n, p)
	hashes := OptimalHashFunctions(p)

	return Bloom{
		Hashes:  hashes,
		Buckets: buckets,
		m:       make([]bool, buckets),
	}
}

// EstimateSize approximates the number of elements the filter has seen.
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

// Add adds a []byte to the bloom filter.
func (b Bloom) Add(data []byte) {
	hashes := hh.GetHashes(b.Hashes, data)
	for _, h := range hashes {
		b.m[h%b.Buckets] = true
	}
}

// Contains checks if the filter has seen the []byte array before
func (b Bloom) Contains(data []byte) bool {
	found := true
	hashes := hh.GetHashes(b.Hashes, data)
	for _, h := range hashes {
		if !b.m[h%b.Buckets] {
			found = false
		}
	}
	return found
}

// Concat concatinates two Bloom filters given that they have the same number of buckets.
// If the number of hash functions used is different we choose the least number, this can
// have an adverse effect on the false positive threshold.
func Concat(a Bloom, b Bloom) (Bloom, error) {
	if a.Buckets != b.Buckets {
		return Bloom{}, errors.New("unable to concatenate Bloom Filters of different Bucket Sizes")
	}
	min := func(a uint64, b uint64) uint64 {
		if a < b {
			return a
		}
		return b

	}
	hashFuncs := min(a.Hashes, b.Hashes)
	c := Bloom{
		Buckets: a.Buckets,
		Hashes:  hashFuncs,
		m:       make([]bool, a.Buckets),
	}

	for i := uint64(0); i < a.Buckets; i++ {
		if a.m[i] || b.m[i] {
			c.m[i] = true
		}
	}
	return c, nil
}
