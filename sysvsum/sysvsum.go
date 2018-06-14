package sysvsum

import "encoding/binary"
import "hash"

type Sum struct {
	state uint32
}

func New() hash.Hash {
	return &Sum{state:0}
}

func (s *Sum) Reset() {
	s.state = 0
}

func (s *Sum) Size() int {
	return 2
}

func (s *Sum) BlockSize() int {
	return 256
}

func (s *Sum) Sum(in []byte) []byte {
	temp := []byte{0,0}
	binary.BigEndian.PutUint16(temp, uint16(s.state % uint32(0xFFFF)))
	return append(in, temp...)
}

func (s *Sum) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state += uint32(v)
	}

	return len(in), nil
}
