package utils

import "encoding/json"

func MarshalJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
