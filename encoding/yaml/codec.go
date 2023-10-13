package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/karlma/soulsbag/encoding"
)

type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (Codec) Decode(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (Codec) String() string {
	return "yaml"
}

func NewEncoder() encoding.Encoder {
	return Codec{}
}

func init() {
	encoding.Register("yaml", NewEncoder)
}
