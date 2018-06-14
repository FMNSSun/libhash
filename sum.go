package libhash

import "encoding/binary"
import "hash"

type Sum32 struct {
	state uint32
}

func New32() hash.Hash {
	return &Sum32{state:0}
}

func (s *Sum32) Reset() {
	s.state = 0
}

func (s *Sum32) Size() int {
	return 4
}

func (s *Sum32) BlockSize() int {
	return 256
}

func (s *Sum32) Sum(in []byte) []byte {
	temp := []byte{0,0,0,0}
	binary.BigEndian.PutUint32(temp, s.state)
	return append(in, temp...)
}

func (s *Sum32) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state += uint32(v)
	}

	return len(in), nil
}

type Sum64 struct {
	state uint64
}

func New64() hash.Hash {
	return &Sum64{state:0}
}

func (s *Sum64) Reset() {
	s.state = 0
}

func (s *Sum64) Size() int {
	return 8
}

func (s *Sum64) BlockSize() int {
	return 256
}

func (s *Sum64) Sum(in []byte) []byte {
	temp := []byte{0,0,0,0,0,0,0,0}
	binary.BigEndian.PutUint64(temp, s.state)
	return append(in, temp...)
}

func (s *Sum64) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state += uint64(v)
	}

	return len(in), nil
}

type Sum16 struct {
	state uint16
}

func New16() hash.Hash {
	return &Sum16{state:0}
}

func (s *Sum16) Reset() {
	s.state = 0
}

func (s *Sum16) Size() int {
	return 2
}

func (s *Sum16) BlockSize() int {
	return 256
}

func (s *Sum16) Sum(in []byte) []byte {
	temp := []byte{0,0,0,0}
	binary.BigEndian.PutUint16(temp, s.state)
	return append(in, temp...)
}

func (s *Sum16) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state += uint16(v)
	}

	return len(in), nil
}

type Sum8 struct {
	state uint8
}

func New8() hash.Hash {
	return &Sum8{state:0}
}

func (s *Sum8) Reset() {
	s.state = 0
}

func (s *Sum8) Size() int {
	return 1
}

func (s *Sum8) BlockSize() int {
	return 256
}

func (s *Sum8) Sum(in []byte) []byte {
	temp := []byte{s.state}
	return append(in, temp...)
}

func (s *Sum8) Write(in []byte) (int, error) {
	for _, v := range in {
		s.state += uint8(v)
	}

	return len(in), nil
}
