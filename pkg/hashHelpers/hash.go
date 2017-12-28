package hashHelpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"hash/crc64"
	"hash/fnv"
)

// TODO this needs to be extended to support values for k larger than 11. This can be done in multiples of 11 by feeding the output from last round into the next. We can also take every 3rd bit and construct new hashes that way.
func GetHashes(k uint64, data []byte) []uint64 {
	sums := []uint64{}
	data = Pad(data, 4)

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

func Pad(b []byte, length int) []byte {
	remaining := length - len(b)

	for i := 0; i < remaining; i++ {
		b = append(b, byte(0))
	}

	return b
}
