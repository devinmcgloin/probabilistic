package minhash

import (
	"math"
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestJaccard(t *testing.T) {

	mh := New(0.05)
	a := generator.RandomSimilarStrings(5000, 3)
	b := generator.RandomSimilarStrings(5000, 3)

	naive := NaiveJaccard(a, b)

	estimate := mh.Estimate(toInterface(a), toInterface(b))
	if naive == 0 {
		t.Error("naive calculated zero overlap")
	}
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
