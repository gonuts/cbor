package cbor

import (
	"io"

	"github.com/ugorji/go/codec"
)

// An Encoder writes CBOR objects to an output stream.
type Encoder struct {
	enc *codec.Encoder
}

// Encode writes the CBOR encoding of v to the stream.
func (enc *Encoder) Encode(v interface{}) error {
	return enc.enc.Encode(v)
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		enc: codec.NewEncoder(w, new(codec.CborHandle)),
	}
}
