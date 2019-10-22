package zero

import (
	"bytes"
	"errors"
	"io"
)

var zero = []byte{0xE2, 0x80, 0x8B} // (U+200B) 0xE2 0x80 0x8B
var one = []byte{0xE2, 0x80, 0x8C}  // (U+200C) 0xE2 0x80 0x8C

func encode(b byte) []byte {
	if b == 0 {
		return zero
	}

	return one
}

// every 3 bytes represent 1 bit
// require 24 bytes to make up 1 byte
// NOTE p must be 24
func decode(p []byte) byte {
	var result byte

	if len(p) != 24 {
		panic("slice must have 24 items")
	}

	for i := 0; i < 24; i += 3 {
		b := p[i : i+3]
		if bytes.Compare(b, one) == 0 {
			// if b is one, then we shift the bytes by index i
			// and or with result
			result |= (0b10000000 >> (i / 3))
		}
		// in case of zero, we don't do anything as result value init with all zeroes
	}

	return result
}

// Encoder is a struct which responsible for transforming
// raw bytes to given zero width chars
type Encoder struct {
	w io.Writer
}

func (se *Encoder) Write(p []byte) (int, error) {
	var n byte

	// This is a very inefficient as it produce 24 bytes per 1 byte, or in other word, 3 bytes per 1 bit
	for _, b := range p {
		n = b & 0b10000000
		se.w.Write(encode(n))
		n = b & 0b01000000
		se.w.Write(encode(n))
		n = b & 0b00100000
		se.w.Write(encode(n))
		n = b & 0b00010000
		se.w.Write(encode(n))
		n = b & 0b00001000
		se.w.Write(encode(n))
		n = b & 0b00000100
		se.w.Write(encode(n))
		n = b & 0b00000010
		se.w.Write(encode(n))
		n = b & 0b00000001
		se.w.Write(encode(n))
	}

	return len(p), nil
}

// NewEncoder convert given io.Writer to zero encoder
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

// Decoder is responsible for transforming zero width chars encoding to original raw bytes.
type Decoder struct {
	r      io.Reader
	buffer []byte
}

func (se *Decoder) Read(p []byte) (int, error) {
	n, err := se.r.Read(se.buffer)
	if err != nil {
		return 0, err
	}

	if n%3 != 0 {
		return 0, errors.New("it's not zero format")
	}

	j := 0

	for i := 0; j < len(p) && i < n; i += 24 {
		p[j] = decode(se.buffer[i : i+24])
		j++
	}

	return j, nil
}

// NewDecoder convert given io.Reader to zero decoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:      r,
		buffer: make([]byte, 240),
	}
}
