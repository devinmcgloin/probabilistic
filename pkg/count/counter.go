package count

import (
	"log"
	"math"

	"github.com/devinmcgloin/probabilistic/pkg/hashHelpers"
)

type Sketch struct {
	b [][]int
	w uint64
	d uint64
}

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

func (s Sketch) Increment(x []byte) {
	hashes := hashHelpers.GetHashes(s.d, x)
	for i, hash := range hashes {
		s.b[i][hash%s.w] += 1
	}
}

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
