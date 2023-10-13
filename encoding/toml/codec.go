package toml

import (
	"github.com/pelletier/go-toml/v2"

	"github.com/karlma/soulsbag/encoding"
)

type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (Codec) Decode(b []byte, v interface{}) error {
	return toml.Unmarshal(b, v)
}

func (Codec) String() string {
	return "toml"
}

func NewEncoder() encoding.Encoder {
	return Codec{}
}

func init() {
	encoding.Register("toml", NewEncoder)
}
