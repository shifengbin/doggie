package yml

import "gopkg.in/yaml.v3"

type YMLDecoder struct{}

func (YMLDecoder) Decode(data []byte, m interface{}) error {
	return yaml.Unmarshal(data, m)
}
