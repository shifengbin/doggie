package json

import "encoding/json"

type JSONDecoder struct{}

func (JSONDecoder) Decode(data []byte, m interface{}) error {
	return json.Unmarshal(data, m)
}
