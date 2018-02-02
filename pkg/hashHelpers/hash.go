package hashHelpers

// GetHashes returns uint64 hashes for a variety of hash functions. Currently k is limited to 11
// TODO this needs to be extended to support values for k larger than 11. This can be done in multiples of 11 by feeding the output from last round into the next. We can also take every 3rd bit and construct new hashes that way.
func GetHashes(k uint64, data []byte) []uint64 {
	hashes := []uint64{}
	for i := uint64(0); i < k; i++ {
		hashes = append(hashes, FNVBias(data, i))
	}
	return hashes
}

// Pad takes a byte array and right pads it with 0
func Pad(b []byte, length int) []byte {
	remaining := length - len(b)

	for i := 0; i < remaining; i++ {
		b = append(b, byte(0))
	}

	return b
}

// FNVHash with bias. Allows constructing unlimited number of hashes.
// see: http://stevehanov.ca/blog/index.php?id=119
// see: http://isthe.com/chongo/tech/comp/fnv/
func FNVBias(b []byte, bias uint64) uint64 {
	var hash uint64
	if bias != 0 {
		hash = bias
	} else {
		hash = 0x01000193
	}

	for _, c := range b {
		hash *= 0x01000193
		hash ^= uint64(c)
		hash &= 0xffffffff
	}

	return hash
}
