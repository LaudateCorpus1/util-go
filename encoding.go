package util

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

func MarshalData(data map[string]interface{}, format string) []byte {
	var marshaledData []byte

	switch format {
	case "json":
		marshaledData, _ = json.Marshal(data)
	case "json-pretty":
		marshaledData, _ = json.MarshalIndent(data, "", "    ")
	case "yaml":
		marshaledData, _ = yaml.Marshal(data)
	}

	return marshaledData
}

func UnmarshalData(byteData []byte, format string) map[string]interface{} {
	var data map[string]interface{}

	switch format {
	case "json":
		json.Unmarshal(byteData, &data)
	case "yaml":
		yaml.Unmarshal(byteData, &data)
	}

	return data
}
