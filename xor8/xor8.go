package xor8

import "hash"

type Xor8 struct {
	state uint8
}

func New() hash.Hash {
	return &Xor8{state:0}
}

func (s *Xor8) Reset() {
	s.state = 0
}

func (s *Xor8) Size() int {
	return 1
}

func (s *Xor8) BlockSize() int {
	return 256
}

func (s *Xor8) Sum(in []byte) []byte {
	temp := []byte{s.state}
	return append(in, temp...)
}

func (s *Xor8) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state ^= v
	}

	return len(in), nil
}
