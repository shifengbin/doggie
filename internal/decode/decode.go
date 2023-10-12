package decode

import (
	"doggie/internal/decode/json"
	"doggie/internal/decode/toml"
	"doggie/internal/decode/yml"
)

type Decoder interface {
	Decode([]byte, interface{}) error
}

var decodes = map[string]Decoder{}

func Register(name string, decoder Decoder) {
	if decoder == nil {
		panic("decoder is nil")
	}

	decodes[name] = decoder
}

func Get(name string) Decoder {
	return decodes[name]
}

func init() {
	Register("json", json.JSONDecoder{})
	Register("yml", yml.YMLDecoder{})
	Register("yaml", yml.YMLDecoder{})
	Register("toml", toml.TOMLDecoder{})
}
