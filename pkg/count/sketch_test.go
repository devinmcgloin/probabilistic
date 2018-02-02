package sketch

import (
	"math"
	"testing"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

func TestCardinality(t *testing.T) {
	b := New(0.001, 0.001)

	url := []byte("https://devinmcgloin.com")
	b.Increment(url)
	if b.Count(url) == 0 {
		t.Errorf("expected estimate count to be %d got %d\n", 1, b.Count(url))
	}

	items := generator.RandomStrings(500)
	for _, v := range items {
		b.Increment([]byte(v))
	}

	for _, str := range items {
		if b.Count([]byte(str)) < 1 {
			t.Errorf("expected estimate count to be %d got %d\n", 1, b.Count([]byte(str)))
		}
	}
}

func TestFalseNegatives(t *testing.T) {
	b := New(0.001, 0.001)
	items := generator.RandomStrings(50000)
	for _, s := range items {
		b.Increment([]byte(s))
	}
	for _, s := range items {
		if b.Count([]byte(s)) < 1 {
			t.Error()
		}
	}
}

func TestLowerBound(t *testing.T) {
	lowerBound := 0.001
	b := New(lowerBound, 0.001)
	members := generator.RandomStrings(3000)
	falsePositives := generator.RandomStrings(4000)

	for _, m := range members {
		b.Increment([]byte(m))
	}

	incorrectCount := 0.0
	for _, f := range falsePositives {
		if b.Count([]byte(f)) > 0 {
			incorrectCount++
		}
	}

	actual := incorrectCount / 4000
	if math.Abs(actual-lowerBound) > 0.1 {
		t.Errorf("Expected lower bound exceeded. Expected %f Actual %f\n", lowerBound, actual)
	}
}
