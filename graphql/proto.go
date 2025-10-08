package graphql

import (
	"encoding/json"
	"fmt"
)

func ProtoKey(message interface{}) string {
	bytes, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("error marshalling message %f", err)
	}
	return string(bytes)
}
