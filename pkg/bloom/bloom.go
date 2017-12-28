package bloom

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash/crc64"
	"hash/fnv"
	"math"
)

type Bloom struct {
	Hashes  uint64
	Buckets uint64
	m       map[uint64]bool
}

// TODO this needs to be extended to support values for k larger than 11. This can be done in multiples of 11 by feeding the output from last round into the next. We can also take every 3rd bit and construct new hashes that way.
func getHashes(k uint64, data []byte) []uint64 {
	sums := []uint64{}

	md5Bytes := md5.Sum(data)
	sums = append(sums, binary.LittleEndian.Uint64(md5Bytes[:]))
	sums = append(sums, binary.BigEndian.Uint64(md5Bytes[:]))
	sha1Bytes := sha1.Sum(data)
	sums = append(sums, binary.LittleEndian.Uint64(sha1Bytes[:]))
	sums = append(sums, binary.BigEndian.Uint64(sha1Bytes[:]))
	sha256Bytes := sha256.Sum256(data)
	sums = append(sums, binary.LittleEndian.Uint64(sha256Bytes[:]))
	sums = append(sums, binary.BigEndian.Uint64(sha256Bytes[:]))
	sha512Bytes := sha512.Sum512(data)
	sums = append(sums, binary.LittleEndian.Uint64(sha512Bytes[:]))
	sums = append(sums, binary.BigEndian.Uint64(sha512Bytes[:]))

	fnvHash := fnv.New32()
	fnvBytes := fnvHash.Sum(data)
	sums = append(sums, binary.LittleEndian.Uint64(fnvBytes[:]))
	sums = append(sums, binary.BigEndian.Uint64(fnvBytes[:]))

	crc64Table := crc64.MakeTable(0xC96C5795D7870F42)
	crc64Int := crc64.Checksum(data, crc64Table)
	sums = append(sums, crc64Int)
	return sums[:k]
}

func OptimalBuckets(n uint64, p float64) uint64 {
	return uint64(-(float64(n) * math.Log(p)) / (math.Log(2) * math.Log(2)))
}

// See https://en.wikipedia.org/wiki/Bloom_filter#Optimal_number_of_hash_functions
func OptimalHashFunctions(p float64) uint64 {
	return uint64(-1.44 * math.Log2(p))
}

func New(n uint64, p float64) Bloom {
	buckets := OptimalBuckets(n, p)
	hashes := OptimalHashFunctions(p)

	return Bloom{
		Hashes:  hashes,
		Buckets: buckets,
		m:       make(map[uint64]bool, buckets),
	}
}

func (b Bloom) EstimateSize() float64 {
	x := 0.0
	for _, v := range b.m {
		if v {
			x++
		}
	}
	return -(float64(b.Buckets) / float64(b.Hashes)) *
		math.Log(1.0-(x/float64(b.Buckets)))
}

func (b Bloom) Add(data []byte) {
	hashes := getHashes(b.Hashes, data)
	for _, h := range hashes {
		b.m[h%b.Buckets] = true
	}
}

func (b Bloom) Contains(data []byte) bool {
	found := true
	hashes := getHashes(b.Hashes, data)
	for _, h := range hashes {
		if !b.m[h%b.Buckets] {
			found = false
		}
	}
	return found
}
