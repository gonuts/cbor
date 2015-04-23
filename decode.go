package cbor

import (
	"io"

	"github.com/ugorji/go/codec"
)

// Decoder reads and decodes CBOR objects from an input stream.
type Decoder struct {
	dec *codec.Decoder
}

// Decode reads the next CBOR-encoded value from its input and
// stores it in the value pointed to by v.
func (dec *Decoder) Decode(v interface{}) error {
	return dec.dec.Decode(v)
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		dec: codec.NewDecoder(r, new(codec.CborHandle)),
	}
}
