package minhash

import (
	"log"
	"math"
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestJaccard(t *testing.T) {

	mh := New(0.05)
	log.Printf("hashes: %d ErrorRate: %f\n", mh.k, mh.ErrorRate())
	a := generator.RandomStrings(5000)
	b := generator.RandomStrings(5000)

	naive := NaiveJaccard(a, b)

	estimate := mh.Estimate(toInterface(a), toInterface(b))
	if math.Abs(naive-estimate) > mh.ErrorRate() {
		t.Error()
	}
}

func toInterface(a []string) []interface{} {
	new := make([]interface{}, len(a))
	for i, v := range a {
		new[i] = v
	}
	return new
}
