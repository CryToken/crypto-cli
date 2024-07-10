package sha4

import (
	"encoding/binary"
	"math/bits"
)

const (
	chunkSize = 64
	hashSize  = 32
)

var k = [64]uint32{
	0x448a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
	0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
	0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
	0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
	0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
	0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
	0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
	0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
	0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

type sha4 struct {
	h   [8]uint32
	x   [chunkSize]byte
	nx  int
	len uint64
}

func (s *sha4) Reset() {
	s.h = [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}
	s.nx = 0
	s.len = 0
}

func New() *sha4 {
	s := new(sha4)
	s.Reset()
	return s
}
func (s *sha4) Hash(p []byte) []byte {
	s.Write(p)
	return s.Sum(nil)
}

func (s *sha4) MultiHash(data []byte, iter int) []byte {
	result := data
	if iter <= 1 {
		return s.Hash(result)
	}
	for i := 0; i < iter; i++ {
		result = s.Hash(result)
	}
	return result

}

func (s *sha4) Write(p []byte) (nn int, err error) {
	nn = len(p)
	s.len += uint64(nn)
	if s.nx > 0 {
		n := copy(s.x[s.nx:], p)
		s.nx += n
		if s.nx == chunkSize {
			block(s, s.x[:])
			s.nx = 0
		}
		p = p[n:]
	}
	if len(p) >= chunkSize {
		n := len(p) &^ (chunkSize - 1)
		block(s, p[:n])
		p = p[n:]
	}
	if len(p) > 0 {
		s.nx = copy(s.x[:], p)
	}
	return
}

func (s *sha4) Sum(in []byte) []byte {
	len := s.len
	var tmp [64]byte
	tmp[0] = 0x80
	if len%64 < 56 {
		s.Write(tmp[0 : 56-len%64])
	} else {
		s.Write(tmp[0 : 64+56-len%64])
	}
	len <<= 3
	binary.BigEndian.PutUint64(tmp[:], len)
	s.Write(tmp[:8])

	var digest [hashSize]byte
	for i, v := range s.h {
		binary.BigEndian.PutUint32(digest[i*4:], v)
	}
	return append(in, digest[:]...)
}

func block(s *sha4, p []byte) {
	var w [64]uint32
	for i := 0; i < 16; i++ {
		w[i] = binary.BigEndian.Uint32(p[4*i:])
	}
	for i := 16; i < 64; i++ {
		w[i] = w[i-16] + (bits.RotateLeft32(w[i-15], -7) ^ bits.RotateLeft32(w[i-15], -18) ^ (w[i-15] >> 3)) +
			w[i-7] + (bits.RotateLeft32(w[i-2], -17) ^ bits.RotateLeft32(w[i-2], -19) ^ (w[i-2] >> 10))
	}

	a, b, c, d, e, f, g, h := s.h[0], s.h[1], s.h[2], s.h[3], s.h[4], s.h[5], s.h[6], s.h[7]
	for i := 0; i < 64; i++ {
		t1 := h + (bits.RotateLeft32(e, -6) ^ bits.RotateLeft32(e, -11) ^ bits.RotateLeft32(e, -25)) + ((e & f) ^ (^e & g)) + k[i] + w[i]
		t2 := (bits.RotateLeft32(a, -2) ^ bits.RotateLeft32(a, -13) ^ bits.RotateLeft32(a, -22)) + ((a & b) ^ (a & c) ^ (b & c))
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}
	s.h[0] += a
	s.h[1] += b
	s.h[2] += c
	s.h[3] += d
	s.h[4] += e
	s.h[5] += f
	s.h[6] += g
	s.h[7] += h
}
