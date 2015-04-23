package cbor_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/gonuts/cbor"
)

func TestCBOR(t *testing.T) {
	for _, table := range []struct {
		data interface{}
		want []byte
	}{
		{
			data: int(42),
			want: []byte{24, 42},
		},
		{
			data: float64(42),
			want: []byte{251, 64, 69, 0, 0, 0, 0, 0, 0},
		},
		{
			data: [2]int64{42, 666},
			want: []byte{130, 24, 42, 25, 2, 154},
		},
		{
			data: [2]float64{42, 666},
			want: []byte{
				130,
				251, 64, 69, 0, 0, 0, 0, 0, 0,
				251, 64, 132, 208, 0, 0, 0, 0, 0,
			},
		},
		{
			data: "boo",
			want: []byte{99, 98, 111, 111},
		},
		{
			data: struct{}{},
			want: []byte{160},
		},
		{
			data: struct {
				X int64
				Y float64
				Z string
			}{X: 42, Y: 42, Z: "boo"},
			want: []byte{
				163,
				97, 88, 24, 42,
				97, 89, 251, 64, 69, 0, 0, 0, 0, 0, 0,
				97, 90, 99, 98, 111, 111,
			},
		},
		{
			data: []int64{1, 2, 3},
			want: []byte{131, 1, 2, 3},
		},
		{
			data: []float64{1, 2, 3},
			want: []byte{
				131,
				251, 63, 240, 0, 0, 0, 0, 0, 0,
				251, 64, 0, 0, 0, 0, 0, 0, 0,
				251, 64, 8, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			data: map[string]int{"one": 1, "two": 2},
			want: []byte{
				162,
				99, 111, 110, 101, 1,
				99, 116, 119, 111, 2,
			},
		},
	} {
		buf := new(bytes.Buffer)
		enc := cbor.NewEncoder(buf)
		err := enc.Encode(table.data)
		if err != nil {
			t.Errorf("could not encode [%#v]: %v\n", table.data, err)
			continue
		}

		if !reflect.DeepEqual(buf.Bytes(), table.want) {
			t.Errorf("wrong encoded value for [%#v]:\ngot= %v\nwant=%v\n",
				table.data,
				buf.Bytes(),
				table.want,
			)
			continue
		}

		val := reflect.New(reflect.ValueOf(table.data).Type())
		dec := cbor.NewDecoder(buf)
		err = dec.Decode(val.Interface())
		if err != nil {
			t.Errorf("could not decode [%#v]: %v\n", table.data, err)
			continue
		}

		if !reflect.DeepEqual(val.Elem().Interface(), table.data) {
			t.Errorf("wrong decoded value for:\ngot= %v\nwant=%v\n",
				val.Elem().Interface(),
				table.data,
			)
		}
	}
}
