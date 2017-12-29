package bloom

import (
	"math"
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestCardinality(t *testing.T) {
	b := New(10, 0.01)
	if b.EstimateSize() != 0 {
		t.Error()
	}

	b.Add([]byte("https://devinmcgloin.com"))
	if b.EstimateSize() == 0 {
		t.Error()
	}
}

func TestFalseNegatives(t *testing.T) {
	b := New(500000, 0.01)
	items := generator.RandomStrings(50000)
	for _, s := range items {
		b.Add([]byte(s))
	}
	for _, s := range items {
		if !b.Contains([]byte(s)) {
			t.Error()
		}
	}
}

func TestLowerBound(t *testing.T) {
	lowerBound := 0.01
	b := New(50000, lowerBound)
	members := generator.RandomStrings(50000)
	falsePositives := generator.RandomStrings(25000)

	for _, m := range members {
		b.Add([]byte(m))
	}

	incorrectCount := 0.0
	for _, f := range falsePositives {
		if b.Contains([]byte(f)) {
			incorrectCount += 1
		}
	}

	actual := incorrectCount / 25000.0
	if math.Abs(actual-lowerBound) > 0.01 {
		t.Errorf("Expected lower bound exceeded. Expected %f Actual %f\n", lowerBound, actual)
	}

}
