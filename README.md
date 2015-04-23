cbor
====

`cbor` is a thin wrapper around [cbor](https://github.com/ugorji/go/codec) and
exposes a familiar `Encoder`/`Decoder` API (a la `json`.)

## Installation

```sh
$ go get github.com/gonuts/cbor
```

## Example

```go
package main

import (
	"bytes"
	"fmt"

	"github.com/gonuts/cbor"
)

func main() {
	buf := new(bytes.Buffer)
	enc := cbor.NewEncoder(buf)
	err := enc.Encode(42)
	if err != nil {
		panic(err)
	}

	val := 0
	dec := cbor.NewDecoder(buf)
	err = dec.Decode(&val)
	if err != nil {
		panic(err)
	}

	fmt.Printf("val=%#v\n", val)
}
```

