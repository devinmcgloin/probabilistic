package minhash

import (
	"bytes"
	"encoding/gob"
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hashHelpers"
)

type MinHash struct {
	k uint64
}

func New(eps float64) MinHash {
	mh := MinHash{}
	mh.k = mh.OptimalHashes(eps)
	return mh
}

func (mh MinHash) OptimalHashes(eps float64) uint64 {
	return uint64(1 / eps * eps)
}

func (mh MinHash) ErrorRate() float64 {
	return 1 / math.Sqrt(float64(mh.k))
}

func (mh MinHash) Estimate(a []interface{}, b []interface{}) float64 {
	y := 0.0
	for i := uint64(0); i < mh.k; i++ {
		if hashMin(a, i) == hashMin(b, i) {
			y++
		}
	}
	return y / float64(mh.k)
}

func hashMin(a []interface{}, bias uint64) uint64 {
	var min uint64
	min = math.MaxUint64
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

func getBytes(a interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(a)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

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
