package utils

import "encoding/json"

func ToJson(form any) string {
	// Thanks to Eson.ninja and sunalwayskonws
	bytes, err := json.Marshal(form)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ToStruct(form any, str string) {
	err := json.Unmarshal([]byte(str), form)
	if err != nil {
		panic(err)
	}
}
