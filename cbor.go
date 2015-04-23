// Package cbor implements encoding and decoding of CBOR objects as defined in
// RFC 7049.
package cbor

import "bytes"

// Marshal returns the CBOR encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	w := new(bytes.Buffer)
	enc := NewEncoder(w)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// Unmarshal parses the CBOR-encoded data and
// stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	r := bytes.NewReader(data)
	dec := NewDecoder(r)
	return dec.Decode(v)
}
