package bloom

import (
	"log"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/devinmcgloin/probabilistic/pkg/generator"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestCardinality(t *testing.T) {
	b := New(5000, 0.01)
	if b.EstimateSize() != 0 {
		t.Error()
	}

	b.Add([]byte("https://devinmcgloin.com"))
	if b.EstimateSize() == 0 {
		t.Errorf("expected estimate size to be %d got %f", 0, b.EstimateSize())
	}

	items := generator.RandomStrings(500)
	for _, v := range items {
		b.Add([]byte(v))
	}

	if math.Abs(b.EstimateSize()-501) > 10 {
		t.Errorf("expected estimate size to be %d got %f", 501, b.EstimateSize())
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
			incorrectCount++
		}
	}

	actual := incorrectCount / 25000.0
	if math.Abs(actual-lowerBound) > 0.01 {
		t.Errorf("Expected lower bound exceeded. Expected %f Actual %f\n", lowerBound, actual)
	}
}

func TestConcat(t *testing.T) {

	lowerBound := 0.01
	a := New(50000, lowerBound)
	b := New(50000, lowerBound)

	members := generator.RandomStrings(30000)
	falsePositives := generator.RandomStrings(25000)

	for _, m := range members {
		if r.NormFloat64() <= 0.5 {
			a.Add([]byte(m))
		} else {
			b.Add([]byte(m))
		}
	}

	c, err := Concat(a, b)
	if err != nil {
		log.Println(err)
	}

	incorrectCount := 0.0
	for _, f := range falsePositives {
		if c.Contains([]byte(f)) {
			incorrectCount++
		}
	}

	actual := incorrectCount / 25000.0
	if math.Abs(actual-lowerBound) > 0.01 {
		t.Errorf("Expected lower bound exceeded under concat. Expected %f Actual %f\n", lowerBound, actual)
	}

}
