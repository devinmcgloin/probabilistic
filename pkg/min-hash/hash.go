package minhash

import (
	"bytes"
	"encoding/gob"
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hashHelpers"
)

// MinHash includes all fields necessary to calculate the min hash between two sets.
type MinHash struct {
	k uint64
}

// New creates a new hash with k set to the number of hashes that
// ensures an error rate of eps.
func New(eps float64) MinHash {
	mh := MinHash{}
	mh.k = mh.OptimalHashes(eps)
	return mh
}

// OptimalHashes calculates the optimal k for a given error rate
func (mh MinHash) OptimalHashes(eps float64) uint64 {
	return uint64(1 / (eps * eps))
}

// ErrorRate is the inverse of OptimalHashes and returns the current expected error rate
func (mh MinHash) ErrorRate() float64 {
	return 1 / math.Sqrt(float64(mh.k))
}

// Estimate takes two sets and returns an estimation of the Jaccard similarity with in
// an error rate eps of the true Jaccard Similarity.
func (mh MinHash) Estimate(a []interface{}, b []interface{}) float64 {
	y := 0.0
	for i := uint64(0); i < mh.k; i++ {
		if hashMin(a, i) == hashMin(b, i) {
			y++
		}
	}
	return y / float64(mh.k)
}

// hashMin calculates the min hash for a given an FNVHash with bias.
func hashMin(a []interface{}, bias uint64) uint64 {
	var min uint64 = math.MaxUint64
	for _, s := range a {
		b, err := getBytes(s)
		if err != nil {
			continue
		}
		hash := hashHelpers.FNVBias(b, bias)
		if hash < min {
			min = hash
		}
	}
	return min
}

// getBytes returns the []byte array for an interface.
func getBytes(a interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(a)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// NaiveJaccard calculates the Jaccard similarity through brute force.
// Used as a comparison in testing Estimate.
func NaiveJaccard(a []string, b []string) float64 {
	union := len(a) + len(b)
	intersection := 0.0

	for _, ai := range a {
		for _, bi := range b {
			if ai == bi {
				intersection++
			}
		}
	}
	return intersection / float64(union)
}
