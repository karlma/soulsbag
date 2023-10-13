package json

import (
	"encoding/json"

	"github.com/karlma/soulsbag/encoding"
)

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for JSON encoding.
type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
	// TODO: expose prefix and indent in the Codec as setting?
	return json.MarshalIndent(v, "", "  ")
}

func (Codec) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func (Codec) String() string {
	return "json"
}

func NewEncoder() encoding.Encoder {
	return Codec{}
}

// import this module from soulsbag
func init() {
	encoding.Register("json", NewEncoder)
}
