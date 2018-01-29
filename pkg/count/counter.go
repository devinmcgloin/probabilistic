package count

import (
	"log"
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hashHelpers"
)

// Sketch represents the data required for min sketch
type Sketch struct {
	b [][]int
	w uint64
	d uint64
}

// New allocates an optimal sketch with eps and confidence parameters.
func New(eps, confidence float64) Sketch {
	w := uint64(math.Ceil(math.E / eps))
	d := uint64(math.Ceil(math.Log(1 / confidence)))

	log.Printf("buckets: %d hashes: %d\n", w, d)
	backing := make([][]int, d)
	for i := range backing {
		backing[i] = make([]int, w)
	}

	return Sketch{
		b: backing,
		w: w,
		d: d,
	}
}

// Increment adds x to the set of seen values
func (s Sketch) Increment(x []byte) {
	hashes := hashHelpers.GetHashes(s.d, x)
	for i, hash := range hashes {
		s.b[i][hash%s.w]++
	}
}

// Count checks how many times x has been seen in the past
func (s Sketch) Count(x []byte) int {
	min := 2 << 32
	hashes := hashHelpers.GetHashes(s.d, x)
	for i, hash := range hashes {
		if min > s.b[i][hash%s.w] {
			min = s.b[i][hash%s.w]
		}
	}
	return min
}

// Stream conusmes a channel of type []byte. This allows the Sketch
// watch a data stream. Only returns when the channel closes. Should
// be run in its own go routine.
func (s Sketch) Stream(c <-chan []byte) {
	for b := range c {
		s.Increment(b)
	}
}
