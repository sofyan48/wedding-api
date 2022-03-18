package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DumpToString interface to string
func DumpToString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		buff := &bytes.Buffer{}
		json.NewEncoder(buff).Encode(v)
		return buff.String()
	}

	return str
}

func DebugPrint(v interface{}) {
	fmt.Println(DumpToString(v))
}
