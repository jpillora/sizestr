package sizestr

import "flag"

//Flag implements the flags.Value interface
type Bytes int64

func Flag(s string) flag.Value {
	var b Bytes
	if err := b.Set(s); err != nil {
		panic(err)
	}
	return &b
}

func (b Bytes) String() string {
	return ToString(int64(b))
}

func (b *Bytes) Set(s string) error {
	bytes, err := Parse(s)
	if err != nil {
		return err
	}
	*b = Bytes(bytes)
	return nil
}
