package toml

import "github.com/pelletier/go-toml/v2"

type TOMLDecoder struct{}

func (TOMLDecoder) Decode(data []byte, m interface{}) error {
	return toml.Unmarshal(data, m)
}
