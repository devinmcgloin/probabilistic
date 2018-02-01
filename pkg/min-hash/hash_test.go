package minhash

import (
	"math"
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestJaccard(t *testing.T) {
	a := generator.RandomStrings(500)
	b := generator.RandomStrings(500)

	naive := NaiveJaccard(a, b)

	mh := New(0.01)
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
